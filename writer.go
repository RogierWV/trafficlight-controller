package main

import (
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
)

func write(msg <-chan bool, state chan StateModCommand, c *websocket.Conn) {
	read := make(chan ControllerState, 1)
	for {
		<-msg
		log.Println("start write")
		state <- StateModCommand{true,nil,read}
		message, err := json.Marshal(<-read)
		if err != nil {
			log.Println(err)
		}
		err = c.WriteMessage((*frameType), message)
		if err != nil {
			log.Println(err)
		}
	}	
}