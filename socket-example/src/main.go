package main

import (
	"socket-example/src/client"
	"socket-example/src/server"
)

func main() {
	server.RunServer()
	client.SendData()
}
