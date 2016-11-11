package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func write(msg <-chan []byte, c *websocket.Conn) {
	for {
		message := <-msg
		err := c.WriteMessage(0x1, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}	
}