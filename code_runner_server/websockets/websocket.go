package websockets

import (
	"bufio"
	"code-runner/code_runner_server/service"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"os"
)

func SourceWebsockets(server *socketio.Server) {
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())

		go func() {
			path := "/logs/" + os.Getenv("PROJECT_ID") + ".log"
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer func(f *os.File) {
				if err := f.Close(); err != nil {
					log.Println("Failed to open file")
				}
			}(f)
			service.WatchLogs(path, func() {
				scanner := bufio.NewScanner(f)
				line := ""
				for scanner.Scan() {
					line += scanner.Text() + "\n"
				}
				if line != "" {
					s.Emit("logs", line)
				}
			})

		}()
		return nil
	})
}
