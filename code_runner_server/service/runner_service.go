package service

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func StartRunner() error {
	if cmd != nil {
		return errors.New("action already running")
	}
	log.Println("Starting runner")
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/start.sh"
	cmd = exec.Command("/bin/sh", scriptPath)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}

func StopRunner() error {
	if cmd != nil {
		log.Println("Stopping process")
		_, err := exec.Command("kill", "-9", string(rune(cmd.Process.Pid))).Output()
		cmd = nil
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
	log.Println("Started installing dependencies")
	scriptPath := os.Getenv("SCRIPTS_PATH") + "/install-dependencies.sh"
	_, err = exec.Command("/bin/sh", scriptPath).Output()
	log.Println("Finished installing dependencies")
	if err != nil {
		return err
	}
	err = StartRunner()
	if err != nil {
		return err
	}
	return nil
}
