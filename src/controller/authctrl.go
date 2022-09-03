package controllers

import (
	"log"
	"net/http"
	"url_manager/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// TODO: Remove
func Authenticate(ctx *gin.Context) {
	s := sessions.Default(ctx)
	id := s.Get("id")
	if id != nil && id != "" && id == ctx.Param("user_id") {
		return
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
}

func Login(ctx *gin.Context) {
	s := sessions.Default(ctx)
	if ok, u := authenticateUser(ctx); ok {
		log.Println(u, s)
		s.Set("id", u.ID)
		s.Save()

		ctx.JSON(http.StatusOK, u)
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{})
}

func Logout(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	ctx.JSON(http.StatusOK, gin.H{})
}

func Register(ctx *gin.Context) {

}

func UnRegister(ctx *gin.Context) {

}

// TODO: Remove
func authenticateUser(ctx *gin.Context) (bool, *model.User) {
	var credential model.User
	ctx.BindJSON(&credential)
	log.Println(credential)
	repo := getUserRepo()
	u := repo.GetByUserId(credential.UserID)
	if u.Authenticate(&credential) {
		return true, &u
	} else {
		return false, nil
	}
}