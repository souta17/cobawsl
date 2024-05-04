package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/souta17/go/configs"
	"github.com/souta17/go/routers"
)

func main() {
	configs.LoadConfig()
	configs.ConnectDb()

	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	api := r.Group("/api")
	routers.AuthRouter(api)
	r.Run(fmt.Sprintf(":%v", configs.ENV.PORT))
}
