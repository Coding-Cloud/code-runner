package main

import (
	"code-runner/code_runner_server"
	"code-runner/code_runner_server/service"
)

func main() {
	service.StartRunner()
	code_runner_server.StartServer()
}
