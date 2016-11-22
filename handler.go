package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true},
}

func HandleWS(c *websocket.Conn) {

	msgChannel := make(chan []byte, 1)
	outputChannel := make(chan bool, 1)
	stateChannel := make(chan StateModCommand, 10)

	go manage_state(stateChannel)
	go process_simstate(msgChannel, outputChannel, stateChannel)
	go write(outputChannel, stateChannel, c)

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		if mt != (*frameType) {
			log.Println("read: incorrect frame type")
			break
		}
		// log.Printf("recv: %s", message)

		msgChannel <- message
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	log.Println("connected")
	HandleWS(c)
}
