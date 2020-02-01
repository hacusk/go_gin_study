package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context) {
	c.HTML(http.StatusMovedPermanently, "error.html", gin.H {
		"title": "エラー！！！",
		"error": "エラー！！！",
	})
}