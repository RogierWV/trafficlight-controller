package main

import (
	"encoding/json"
	"log"
)

func process_simstate(msg <-chan []byte, out chan<- bool, state chan<- StateModCommand, timer chan<- int) {
	var simState SimulatorState

	retChan := make (chan ControllerState)

	for {
		message := <-msg

		err := json.Unmarshal(message, &simState)
		if err != nil {
			log.Println(err)
		}

		green := SimulatorStateSub{-1,0}

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
					// set_all_red(contrState)
						(*contrState).State[green.TrafficLight] = ControllerStateSub{green.TrafficLight,"green"}
					ret <- (*contrState)
				},
				retChan,
			}

			timer <- green.TrafficLight
		}

		// contrState := <-retChan

		
		out <- true
	}
}