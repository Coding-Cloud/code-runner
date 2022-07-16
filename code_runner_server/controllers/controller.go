package controllers

import "github.com/gin-gonic/gin"

func SourceControllers(engine *gin.Engine) {
	engine.POST("/start", startProject)
	engine.POST("/restart", restartProject)
	engine.POST("/dependencies", installDependencies)
	engine.POST("/stop", stopProject)
}
