package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SourceControllers(engine *gin.Engine) {
	engine.GET("/status", func(context *gin.Context) {
		context.Writer.WriteHeader(http.StatusOK)
	})
	engine.POST("/start", startProject)
	engine.POST("/restart", restartProject)
	engine.POST("/dependencies", installDependencies)
}
