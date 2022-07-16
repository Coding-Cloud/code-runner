package main

import (
	"code-runner/code_runner_server/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	controllers.SourceControllers(engine)

	err := engine.Run(":8080")
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
