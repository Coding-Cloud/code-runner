package code_runner_server

import (
	"code-runner/code_runner_server/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func StartServer() {
	engine := gin.Default()

	controllers.SourceControllers(engine)
	port := os.Getenv("RUNNER_PORT")
	err := engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
