package controllers

import (
	"code-runner/code_runner_server/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func startProject(c *gin.Context) {
	err := service.StartRunner()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func restartProject(c *gin.Context) {
	err := service.RestartRunner()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func installDependencies(c *gin.Context) {
	go func() {
		err := service.InstallDependencies()
		if err != nil {
			log.Print(err)
		}
	}()
	c.Writer.WriteHeader(http.StatusOK)
}
