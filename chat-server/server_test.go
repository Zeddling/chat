package main

import (
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

/*
Test if server returns the sent message
*/
func TestWebsocketServer(t *testing.T) {
	//	set up websocket connection
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8000",
		Path:   "/chat",
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("Failed to connect websocket server: %v", err)
	}
	defer conn.Close()

	//	send message to server
	msg := []byte("Hello")
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		t.Fatalf("Failed to send message to Websocket server: %v", err)
	}

	//	Read the server's response
	_, response, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message from websocket server: %v", err)
	}

	if string(response) != string(msg) {
		t.Fatalf("Unexpected response from websocket server. \nExpected: %s; \nGot: %s", msg, response)
	}
}
