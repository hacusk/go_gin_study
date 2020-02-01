package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMenu View Controll "/menu"
func GetMenu(c *gin.Context) {

	menu := []string{"ししゃもっこ", "さくらんぼ"}

	c.HTML(http.StatusOK, "menu.html", gin.H{
		"menu": menu,
	})
}
