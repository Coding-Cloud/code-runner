package service

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func StartRunner() error {
	if cmd == nil {
		log.Print("Already running")
		return errors.New("runner already running")
	}
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/start.sh"
	log.Print("Starting ")
	cmd = exec.Command("/bin/sh", scriptPath)
	if err := cmd.Start(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func StopRunner() error {
	if cmd == nil {
		cmd = nil
		err := cmd.Process.Kill()
		if err != nil {
			log.Print(err)
			return err
		}
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
	err := StopRunner()
	if err != nil {
		return err
	}
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/install-dependencies.sh"
	_, err = exec.Command("/bin/sh", scriptPath).Output()
	if err != nil {
		log.Print(err)
		return err
	}
	err = StartRunner()
	if err != nil {
		return err
	}
	return nil
}
