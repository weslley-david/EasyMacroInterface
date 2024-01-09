package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
)

func main() {
	r := gin.Default()
	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "1" {
			robotgo.KeyTap("q", "alt")
		}
		if id == "2" {
			robotgo.KeyTap("w", "alt")
		}
		if id == "3" {
			robotgo.KeyTap("a", "alt")
		}
		if id == "4" {
			robotgo.KeyTap("s", "alt")
		}
		if id == "5" {
			robotgo.KeyTap("z", "alt")
		}
		if id == "6" {
			robotgo.KeyTap("cmd")
		}

		c.JSON(200, gin.H{
			"message": id,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
