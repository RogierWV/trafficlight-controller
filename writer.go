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
			log.Fatal(err)
		}
		err = c.WriteMessage(0x1, message)
		if err != nil {
			log.Fatal(err)
		}
	}	
}