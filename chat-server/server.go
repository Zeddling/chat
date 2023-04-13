// Handles client configuration for connection
// to the chat server

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct{}

var upgrader = websocket.Upgrader{}

func NewServer() Server {
	return Server{}
}

func (server *Server) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("Error during connection upgradation", err)
		return
	}
	defer conn.Close()

	log.Println("Websocket handler for url /chat has been created")
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing: ", err)
			break
		}
	}
}
