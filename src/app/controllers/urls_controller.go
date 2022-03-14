package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"url_manager/app/models"
	"url_manager/app/models/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowURLs(c *gin.Context) {

}

func ShowURL(c *gin.Context) {

}

func NewURL(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("uid")

	c.HTML(http.StatusOK, "url_new.html", gin.H{"id": id})
}

func CreateURL(c *gin.Context) {
	session := sessions.Default(c)
	id := fmt.Sprintf("%d", session.Get("uid"))

	c.Request.ParseForm()
	title := c.Request.FormValue("title")
	url := c.Request.FormValue("url")

	intid, _ := strconv.Atoi(id)

	urlModel := models.URL{
		Title:  title,
		URL:    url,
		UserID: uint(intid),
	}

	r := repositories.DefaultURLRepositoryImpl{}
	r.Create(urlModel)
	a, _ := r.GetAll()
	fmt.Println(a)
	c.Redirect(302, "/users/"+id)
}

func EditURL(c *gin.Context) {

}

func UpdateURL(c *gin.Context) {

}

func DeleteURL(c *gin.Context) {

}