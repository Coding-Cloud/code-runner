package service

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

func WatchLogs(path string, callback func()) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer func(watcher *fsnotify.Watcher) {
		if err := watcher.Close(); err != nil {
			log.Fatal("Failed to close watcher")
		}
	}(watcher)

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					callback()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watch error:", err)
			}
		}
	}()

	if err := watcher.Add(path); err != nil {
		log.Println("ERROR", err)
	}
	log.Println("Started watching file")
	<-done
}

func ClearLogs() {
	path := "/logs/" + os.Getenv("PROJECT_ID") + ".log"
	if err := os.Remove(path); err != nil {
		return
	}
}
