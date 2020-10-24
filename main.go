package main

import (
	"com.github/asanobm/golang-docker-temp/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.Conn()	
	defer db.Close()
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message" : "pong",
		})
	})
	r.Run()
}