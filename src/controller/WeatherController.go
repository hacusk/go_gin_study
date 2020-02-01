package controller

import (
	"example.com/go_gin_study/src/service"
	"github.com/gin-gonic/gin"
)

// GetWeatherResponse View Controll "/weather"
func GetWeatherResponse(c *gin.Context) {
	service.GetWeather()
}
