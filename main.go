package main

// Test value:
// {"state":[{"trafficLight":1,"count":1}]}
// {"state":[{"trafficLight":30,"count":10},{"trafficLight":21,"count":1}]}
// Expected return value:
// {"state":[{"trafficLight":1,"status":"green"}]}

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// import _ "net/http/pprof"

var addr = flag.String("addr", "0.0.0.0:3000", "http service address (\"$ADDR:$PORT\"")
var frameType = flag.Int("ft", 0x1, "frame type (1 = text, 2 = binary)")
var redTime = flag.Int("ontrtijd", 2, "ontruimingstijd")
var randomise = flag.Bool("r", false, ¨"randomise loop for not fully compliant clients")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handler)
	fmt.Println("listening on ", *addr)
	log.Println(http.ListenAndServe(*addr, nil))
}
