package service

import (
	"log"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func startRunner() {
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/start.sh"
	cmd := exec.Command("/bin/sh", scriptPath)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func stopRunner() {
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatal(err)
	}
}

func restartRunner() {
	stopRunner()
	startRunner()
}

func installDependencies() {
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/install-dependencies.sh"
	_, err := exec.Command("/bin/sh", scriptPath).Output()
	if err != nil {
		log.Fatal(err)
	}
}
