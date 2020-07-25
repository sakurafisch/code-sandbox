package main

import (
	"socket-example/src/client"
	"socket-example/src/server"
)

func main() {
	go server.RunServer()
	client.SendData()
}
