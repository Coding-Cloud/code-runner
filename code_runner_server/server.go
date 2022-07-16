package code_runner_server

import (
	"code-runner/code_runner_server/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	engine := gin.Default()

	controllers.SourceControllers(engine)

	err := engine.Run(":8181")
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
