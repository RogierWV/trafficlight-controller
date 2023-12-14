package main

import (
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/xeipuuv/gojsonschema"
	// "bytes"
)

func TestWS(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	schema := gojsonschema.NewReferenceLoader("file://./schema.json")
	// log.Println(schema)
	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			doc := gojsonschema.NewStringLoader(string(message))
			// log.Println(doc)
			result, err := gojsonschema.Validate(schema, doc)
			if err != nil {
				panic(err.Error())
			}
			if result.Valid() {
				log.Printf("The document is valid\n")
			} else {
				log.Printf("The document is not valid. see errors :\n")
				for _, desc := range result.Errors() {
					log.Printf("- %s\n", desc)
				}
				log.Fatal("incorrect result")
			}
		}
	}()

	c.WriteMessage(websocket.TextMessage, []byte(`{"state":[{"trafficLight":30,"count":10},{"trafficLight":21,"count":1}]}`))

	time.Sleep(1 * time.Second)

	c.WriteMessage(websocket.TextMessage, []byte(`{"state":[{"trafficLight":30,"count":0},{"trafficLight":21,"count":0}]}`))

	time.Sleep(20 * time.Second)

}
