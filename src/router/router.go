package router

import (
	"example.com/go_gin_study/src/common/handler"
	"example.com/go_gin_study/src/controller"

	"github.com/gin-gonic/gin"
)

// GetRouter hoge
func GetRouter() *gin.Engine {

	router := gin.Default()

	router.Static("/static", "resources/static")
	router.LoadHTMLGlob("resources/templates/*")

	router.GET("/", controller.GetIndex)

	router.POST("/auth", controller.PostAuth)

	router.GET("/menu", controller.GetMenu)

	router.GET("/weather", controller.GetWeatherResponse)

	router.NoRoute(handler.Error)

	return router
}
