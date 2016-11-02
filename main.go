package main

import (
	"flag"
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
)

var addr = flag.String("addr", "0.0.0.0:3000", "http service address")
var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
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