package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// writes to client on input on the msg channel, reads state via state channel
func write(msg <-chan bool, state chan ContrStateModCommand, c *websocket.Conn) {
	log.Println("writer started")
	read := make(chan ControllerState, 1)
	for {
		<-msg
		log.Println("start write")
		state <- ContrStateModCommand{true, nil, read}
		message, err := json.Marshal(<-read)
		if err != nil {
			log.Println(err)
		}
		log.Printf("send: %s", message)
		err = c.WriteMessage((*frameType), message)
		if err != nil {
			log.Println(err)
		}
	}
}
