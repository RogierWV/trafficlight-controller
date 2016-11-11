package main

import (
	"encoding/json"
	"log"
)

func process_simstate(msg <-chan []byte, out chan<- []byte) {
	var simState SimulatorState
	contrState := ControllerState{ State: make ([]ControllerStateSub, 40) }

	for i := 0; i < len(contrState.State); i++ {
		contrState.State[i] = ControllerStateSub{i,"red"}
	}

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
			for i := 0; i < len(contrState.State); i++ {
				if contrState.State[i].Status == "green" {
					contrState.State[i].Status = "orange"
				} else {
					contrState.State[i].Status = "red"
				}
			}
			contrState.State[green.TrafficLight].Status = "green"
		}

		newmessage, err := json.Marshal(&contrState)
		if err != nil {
			log.Fatal(err)
		}

		out <- newmessage
	}
}