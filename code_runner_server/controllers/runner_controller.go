package controllers

import (
	"code-runner/code_runner_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func startProject(c *gin.Context) {
	service.StartRunner()
	c.Writer.WriteHeader(http.StatusOK)
}

func restartProject(c *gin.Context) {
	service.RestartRunner()
	c.Writer.WriteHeader(http.StatusOK)
}

func stopProject(c *gin.Context) {
	service.StopRunner()
	c.Writer.WriteHeader(http.StatusOK)
}

func installDependencies(c *gin.Context) {
	service.InstallDependencies()
	c.Writer.WriteHeader(http.StatusOK)
}
