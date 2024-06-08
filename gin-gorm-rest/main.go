package main

import (
	"gin-mod-rest/config"
	"gin-mod-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Conncect()
	routes.UserRoute(router)
	router.Run(":8080")
}
