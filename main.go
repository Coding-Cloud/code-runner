package main

import (
	"code-runner/code_runner_server/controllers"
	"code-runner/code_runner_server/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	controllers.SourceControllers(engine)
	service.StartRunner()
	defer service.StopRunner()

	err := engine.Run(":8080")
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
