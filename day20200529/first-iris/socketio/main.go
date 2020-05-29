package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kataras/iris"
	"log"
)

func main() {
	app := iris.New()
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(conn socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		conn.Emit("reply", "have " + msg)
	})

	server.OnEvent("/chat", "msg", func(conn socketio.Conn, msg string) string {
		conn.SetContext(msg)
		return "recv" + msg
	})

	server.OnEvent("/", "bye", func(conn socketio.Conn) string {
		last := conn.Context().(string)
		conn.Emit("bye", last)
		conn.Close()
		return last
	})

	server.OnError("/", func(conn socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	server.OnDisconnect("/", func(conn socketio.Conn, s string) {
		fmt.Println("closed", s)
	})

	go server.Serve()
	defer server.Close()

	app.HandleMany("GET POST", "/socket.io/{any:path}", iris.FromStd(server))
	app.HandleDir("/", "./asset")
	app.Listen(":8000", iris.WithoutPathCorrection)
}
