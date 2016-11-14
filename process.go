package main

import (
	"encoding/json"
	"log"
	"fmt"
)

func process_simstate(msg <-chan []byte, out chan<- []byte, state chan<- StateModCommand) {
	var simState SimulatorState

	retChan := make (chan ControllerState)

	for {
		message := <-msg

		err := json.Unmarshal(message, &simState)
		if err != nil {
			log.Fatal(err)
		}

		green := SimulatorStateSub{-1,-1}

		//set a green light
		for _,e := range simState.State {
			if e.Count > green.Count {
				green = e
			}
		}

		if green.TrafficLight != -1 {
			state <- StateModCommand{
				false,
				func(contrState *ControllerState, ret chan<- ControllerState){
					fmt.Println("mod started")
					set_all_red(contrState)
					(*contrState).State[green.TrafficLight].Status = "green"
					ret <- (*contrState)
				},
				retChan,
			}
		}

		contrState := <-retChan

		newmessage, err := json.Marshal(&contrState)
		if err != nil {
			log.Fatal(err)
		}

		out <- newmessage
	}
}