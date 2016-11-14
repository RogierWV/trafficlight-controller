package main

// Test value:
// {"state":[{"trafficLight":1,"count":1}]}
// {"state":[{"trafficLight":30,"count":10},{"trafficLight":21,"count":1}]}
// Expected return value:
// {"state":[{"trafficLight":1,"status":"green"}]}

import (
	"flag"
	"net/http"
	"fmt"
	"log"
)

import _ "net/http/pprof"

var addr = flag.String("addr", "0.0.0.0:3000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handler)
	fmt.Println("listening on ", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}