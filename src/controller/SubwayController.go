package controller

import (
	"example.com/go_gin_study/src/service"
	"github.com/gin-gonic/gin"
)

// GetSubwayResponse View Controll "/subway"
func GetSubwayResponse(c *gin.Context) {
	service.GetSubway()
}
