package main

import (
	"github.com/gin-gonic/gin"
	jsonp "github.com/li-jin-gou/hertz-jsonp"
)

func main() {
	r := gin.Default()
	r.Use(jsonp.JsonP())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8080") // listen and server on 0.0.0.0:8080
}
