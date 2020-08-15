package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello, World!")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatalln(err)
	}
	if server != nil {
		server.Close()
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID)
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) error {
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error: ", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Server at localhost:8000...")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
