package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
)

func main() {
	r := gin.Default()
	//===========================================

	r.LoadHTMLGlob("templates/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"host": "192.168.0.110:8080",
		})
	})

	//===========================================
	r.Static("/assets", "./assets")
	//===========================================
	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Print(id)

		if id == "1" {
			robotgo.KeyTap("q", "alt")
		}
		if id == "2" {
			//exec.Command("opera")
			robotgo.KeyTap("w", "alt")
		}
		if id == "3" {
			robotgo.KeyTap("a", "alt")
		}
		if id == "4" {
			robotgo.KeyTap("s", "alt")
		}
		if id == "5" {
			robotgo.KeyTap("a", "alt")
		}
		if id == "6" {
			robotgo.KeyTap("cmd")
		}
		if id == "7" {
			robotgo.KeyTap("cmd")
		}
		if id == "8" {
			robotgo.KeyTap("cmd")
		}

		c.JSON(200, gin.H{
			"message": id,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
