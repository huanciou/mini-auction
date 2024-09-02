package controllers

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

var (
	wg sync.WaitGroup
)

func SetupSocketIO() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())

		wg.Add(1)
		go func() {
			defer wg.Done()
			returnPrice := GetAuction()
			s.Emit("reply", returnPrice)
		}()

		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("Received message:", msg)
		bid, _ := strconv.Atoi(msg)

		success, price := PostAuction(bid, "John Doe")

		if success {
			server.BroadcastToNamespace("/", "reply", price)
		} else {

			message := "Your bid is lower"
			s.Emit("reply", message)
		}
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Printf("socketio listen error: %s\n", err)
		}
	}()
	// defer server.Close()

	return server
}
