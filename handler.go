package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true},
}

func HandleWS(c *websocket.Conn) {
	var simState SimulatorState
	var contrState ControllerState

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		// LOGIC
		// inDec := json.NewDecoder(bytes.NewReader(message))
		err = json.Unmarshal(message, &simState)
		if err != nil {
			log.Fatal(err)
		}
		contrState = ControllerState{ State: []ControllerStateSub{ ControllerStateSub{TrafficLight: simState.State[0].TrafficLight, Status:"green"} }  }
		message, err = json.Marshal(&contrState)
		if err != nil {
			log.Fatal(err)
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	HandleWS(c)
}
