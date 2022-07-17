package service

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

var isRunning bool

func StartRunner() error {
	if isRunning {
		return errors.New("action already running")
	}
	log.Println("Starting runner")
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/start.sh"
	cmd := exec.Command("/bin/sh", scriptPath)
	if err := cmd.Start(); err != nil {
		return err
	}
	isRunning = true
	return nil
}

func StopRunner() error {
	if isRunning {
		log.Println("Stopping process")
		scriptPath := os.Getenv("SCRIPTS_PATH") + "/stop.sh"
		cmd := exec.Command("/bin/sh", scriptPath)
		if err := cmd.Start(); err != nil {
			log.Print(err)
			return err
		}
		isRunning = false
	}
	return nil
}

func RestartRunner() error {
	err := StopRunner()
	if err != nil {
		return err
	}
	err = StartRunner()
	if err != nil {
		return err
	}
	return nil
}

func InstallDependencies() error {
	if err := StopRunner(); err != nil {
		return err
	}
	log.Println("Started installing dependencies")
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/install-dependencies.sh"
	_, err := exec.Command("/bin/sh", scriptPath).Output()
	log.Println("Finished installing dependencies")
	if err != nil {
		return err
	}
	isRunning = false

	if err := StartRunner(); err != nil {
		return err
	}
	return nil
}
