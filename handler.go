package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handle_ws(c *websocket.Conn) {

	msgChannel := make(chan []byte, 1)
	outputChannel := make(chan bool, 1)
	contrStateChannel := make(chan ContrStateModCommand, 10)
	simStateChannel := make(chan SimStateModCommand, 10)

	go manage_sim_state(simStateChannel)
	go update_sim_state(msgChannel, simStateChannel)
	go manage_controller_state(contrStateChannel)
	go process_simstate(outputChannel, contrStateChannel, simStateChannel)
	go write(outputChannel, contrStateChannel, c)

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
	handle_ws(c)
}
