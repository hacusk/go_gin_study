package main

import (
	"example.com/go_gin_study/src/router"
)

func main() {

	router := router.GetRouter()
	router.Run(":8080")
}
