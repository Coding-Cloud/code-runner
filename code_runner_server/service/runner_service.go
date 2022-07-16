package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func startRunner() {
	fmt.Println("FOO:", os.Getenv("SCRIPTS_PATH"))
	cmd := exec.Command("/bin/sh", "/path/to/file.sh")
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
	_, err := exec.Command("/bin/sh", "/path/to/file.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
}
