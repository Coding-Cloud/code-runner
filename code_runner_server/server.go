package code_runner_server

import (
	"code-runner/code_runner_server/controllers"
	"code-runner/code_runner_server/service"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	engine := gin.Default()

	controllers.SourceControllers(engine)
	err := service.StartRunner()
	if err != nil {
		log.Fatal("unable to start runner on launch")
	}
	defer func() {
		err := service.StopRunner()
		if err != nil {
			log.Fatal("unable to stop runner")
		}
	}()

	err = engine.Run(":8080")
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
