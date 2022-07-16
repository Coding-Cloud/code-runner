package controllers

import (
	"code-runner/code_runner_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func startProject(c *gin.Context) {
	service.StartRunner()
	c.Writer.WriteHeader(http.StatusCreated)
}

func restartProject(c *gin.Context) {
	service.RestartRunner()
	c.Writer.WriteHeader(http.StatusCreated)
}

func stopProject(c *gin.Context) {
	service.StopRunner()
	c.Writer.WriteHeader(201)
}

func installDependencies(c *gin.Context) {
	service.InstallDependencies()
	c.Writer.WriteHeader(201)
}
