package controllers

import "github.com/gin-gonic/gin"

func startProject(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func stopProject(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func installDependencies(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
