package main

import (
	"encoding/json"
	"log"
)

func process_simstate(msg <-chan []byte, out chan<- bool, state chan<- StateModCommand, timer chan<- int) {
	simState := SimulatorState{make([]SimulatorStateSub, 50)}
	for i := 0; i < len(simState.State); i++ {
		simState.State[i] = SimulatorStateSub{i,0}
	}

	// retChan := make (chan ControllerState)

	for {
		message := <-msg

		var tmpSimState SimulatorState

		err := json.Unmarshal(message, &tmpSimState)
		if err != nil {
			log.Println(err)
		}

		for _,e := range tmpSimState.State {
			simState.State[e.TrafficLight] = e
		}

		// green := SimulatorStateSub{-1,0}

		// //set a green light
		// for _,e := range simState.State {
		// 	if e.Count > green.Count {
		// 		green = e
		// 	}
		// }

		// if green.TrafficLight != -1 {
		// 	state <- StateModCommand{
		// 		false,
		// 		func(contrState *ControllerState, ret chan<- ControllerState){
		// 			// set_all_red(contrState)
		// 				(*contrState).State[green.TrafficLight] = ControllerStateSub{green.TrafficLight,"green"}
		// 			ret <- (*contrState)
		// 		},
		// 		retChan,
		// 	}

		// 	timer <- green.TrafficLight
		// }

		// state <- StateModCommand{true,nil,retChan}
		highestTotal := 0
		groupId := -1
		// currState := <-retChan

		for i := 0; i < len(lightGroups); i++ {
			total := 0
			for j := 0; j < len(lightGroups[i]); j++ {
				total += simState.State[lightGroups[i][j]].Count 
			}
			if total > highestTotal {
				highestTotal = total
				groupId = i
			}
		}

		if groupId != -1 {
			state <- StateModCommand {
				false,
				func(contrState *ControllerState, ret chan<- ControllerState){
					for _,e := range lightGroups[groupId] {
						(*contrState).State[e] = ControllerStateSub{e,"green"}
					}
				},
				nil,
			}

		}
		
		out <- true
	}
}