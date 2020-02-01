package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex View Controll "/"
func GetIndex(c *gin.Context) {

	name := "Hello, World!"
	title := "TEST"

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": title,
		"name":  name,
	})
}

// PostAuth Auth Controll
func PostAuth(c *gin.Context) {

	if "test" == c.PostForm("name") {
		c.Redirect(http.StatusMovedPermanently, "/menu")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	return
}
