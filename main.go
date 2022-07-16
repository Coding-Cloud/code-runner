package main

import (
	"code-runner/code_runner_server"
	"code-runner/code_runner_server/service"
	"log"
)

func main() {
	err := service.InstallDependencies()
	if err != nil {
		log.Fatal("unable to start runner on launch", err)
	}
	defer func() {
		err := service.StopRunner()
		if err != nil {
			log.Fatal("unable to stop runner", err)
		}
	}()
	code_runner_server.StartServer()
}
