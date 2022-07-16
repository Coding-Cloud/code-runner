package controllers

import (
	"code-runner/code_runner_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func startProject(c *gin.Context) {
	err := service.StartRunner()
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func restartProject(c *gin.Context) {
	err := service.RestartRunner()
	if err != nil {
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func stopProject(c *gin.Context) {
	err := service.StopRunner()
	if err != nil {
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func installDependencies(c *gin.Context) {
	err := service.InstallDependencies()
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
