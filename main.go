package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run(":5050")

	log.Println("Server is running on port 5050")



}