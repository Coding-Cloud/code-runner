package service

import (
	"fmt"
	"os"
	"os/exec"
)

func startProject() {
	fmt.Println("FOO:", os.Getenv("SCRIPTS_PATH"))
	_, err := exec.Command("/bin/sh", "/path/to/file.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func stopProject() {
	_, err := exec.Command("/bin/sh", "/path/to/file.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func installDependencies() {
	_, err := exec.Command("/bin/sh", "/path/to/file.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
