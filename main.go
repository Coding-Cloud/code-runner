package main

import (
	"code-runner/code_runner_server"
	"code-runner/code_runner_server/service"
	"log"
	"os"
)

func main() {
	projectId := os.Getenv("PROJECT_ID")
	log.Printf("Starting code runner for project: %s\n", projectId)
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
