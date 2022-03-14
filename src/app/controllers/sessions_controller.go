package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"url_manager/app/models/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewSession(c *gin.Context) {
	c.HTML(http.StatusOK, "session/login.html", gin.H{"title": "login"})
}

func CreateSession(c *gin.Context) {
	c.Request.ParseForm()
	name := c.Request.FormValue("userId")
	password := c.Request.FormValue("password")

	var err error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(name, password)

	session := sessions.Default(c)
	// if IsLoggedIn(session) {
	// 	c.JSON(http.StatusOK, gin.H{"err": "already login"})
	// 	return
	// }

	if Authenticate(name, password) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	u, err := repositories.UserRepository{}.GetByName(name)
	Login(session, u.ID)

	c.Redirect(302, "/users/"+strconv.Itoa(int(u.ID)))
}

func IsLoggedIn(s sessions.Session) bool {
	return s.Get("uid") != nil
}

func Authenticate(name string, password string) error {
	u, err := repositories.UserRepository{}.GetByName(name)

	fmt.Println(u, err)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func Login(s sessions.Session, id uint) {
	s.Set("uid", id)
	s.Save()
}

func DestroySession(c *gin.Context) {
	session := sessions.Default(c)
	Logout(session)
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(s sessions.Session) {
	s.Clear()
	s.Save()
}
