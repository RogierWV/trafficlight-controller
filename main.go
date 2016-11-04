package main

// Test value:
// {"state":[{"trafficLight":"1","count":"1"}]}

import (
	"flag"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"encoding/json"
	// "bytes"
)

type ControllerState struct {
	State []ControllerStateSub `json:"state"`
}

type ControllerStateSub struct {
	TrafficLight int `json:"trafficLight"`
	Status string `json:"status"`
}

type SimulatorState struct {
	State []SimulatorStateSub `json:"state"`
} 

type SimulatorStateSub struct {
	TrafficLight int `json:"trafficLight"`
	Count int `json:"count"`
}

var addr = flag.String("addr", "0.0.0.0:3000", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true},
}

func echo(w http.ResponseWriter, r *http.Request) {
	var simState SimulatorState
	var contrState ControllerState

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
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

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)
	fmt.Println("listening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}