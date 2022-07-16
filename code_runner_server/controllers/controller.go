package controllers

import "github.com/gin-gonic/gin"

func SourceControllers(engine *gin.Engine) {
	engine.POST("/", startProject)
	engine.POST("/dependencies", installDependencies)
	engine.DELETE("/", stopProject)
}
