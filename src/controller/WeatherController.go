package controller

import (
	"example.com/go_gin_study/src/service"
	"github.com/gin-gonic/gin"
)

// GetWeatherResponse View Controll "/subway"
func GetWeatherResponse(c *gin.Context) {
	service.GetWeather()
}
