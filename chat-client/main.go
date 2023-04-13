package main

func main() {
	client := NewClient()
	defer client.Conn.Close()

	go client.readAndSend()
	go client.receiveHandler()

	client.MessageHandler()
}
