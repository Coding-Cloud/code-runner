package code_runner_server

import (
	"code-runner/code_runner_server/controllers"
	"code-runner/code_runner_server/websockets"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"os"
)

func StartServer() {
	engine := gin.Default()

	server := socketio.NewServer(nil)
	websockets.SourceWebsockets(server)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer func(server *socketio.Server) {
		err := server.Close()
		if err != nil {
			log.Fatalf("socketio close error: %s\n", err)
		}
	}(server)

	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))

	controllers.SourceControllers(engine)
	port := "8080"
	if value, isPresent := os.LookupEnv("RUNNER_PORT"); isPresent {
		port = value
	}
	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("error %s", err)
	}
}
