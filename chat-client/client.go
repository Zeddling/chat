package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn      *websocket.Conn
	done      chan interface{}
	interrupt chan os.Signal
	msg       chan string
}

func NewClient() Client {
	socketUrl := "ws://localhost:8000/chat"

	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to websocket server: ", err)
	}

	return Client{
		Conn:      conn,
		done:      make(chan interface{}),
		interrupt: make(chan os.Signal),
		msg:       make(chan string),
	}
}

func (client *Client) MessageHandler() {
	//	Main loop for the client
	//	Send relevant packets here
	for {
		select {
		case m := <-client.msg:
			//	Send message to server
			fmt.Println(m)
			err := client.Conn.WriteMessage(websocket.TextMessage, []byte(m))
			if err != nil {
				log.Println("Error during writing to websocket: ", err)
				return
			}
		case <-client.interrupt:
			//	We received a SIGINT (CTRL + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			//	Close our websocket connection
			err := client.Conn.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			)
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-client.done:
				log.Println("Receiver Channel Closed! Exiting...")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting...")
			}
			return
		}
	}
}

func (client *Client) receiveHandler() {
	defer close(client.done)

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
		go client.readAndSend()
	}
}

func (client *Client) readAndSend() {
	fmt.Printf("Me: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Encountered error: ", err)
	}
	client.msg <- line
}
