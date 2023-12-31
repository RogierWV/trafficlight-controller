package main

import (
	"encoding/json"
	"log"
	"time"
)

// Keeps controller state (lights), to be run in a separate goroutine
func manage_controller_state(queue <-chan ContrStateModCommand) {
	time.Sleep(1 * time.Second)
	contrState := ControllerState{State: make([]ControllerStateSub, 50)}
	for i := 0; i < len(contrState.State); i++ {
		if nodes[i] != dummyWL {
			contrState.State[i] = ControllerStateSub{i, "red"}
		}
	}
	for {
		command := <-queue
		if command.ReadOnly {
			go func() { command.Ret <- contrState }()
		} else {
			go command.Modifier(&contrState, command.Ret)
		}
	}
}

// Keeps simulator state (counts), to be run in a separate goroutine
func manage_sim_state(queue <-chan SimStateModCommand) {
	time.Sleep(1 * time.Second)
	simState := SimulatorState{State: make([]SimulatorStateSub, 50)}
	for i := 0; i < len(simState.State); i++ {
		simState.State[i] = SimulatorStateSub{i, 0}
	}

	for {
		command := <-queue
		if command.ReadOnly {
			go func() { command.Ret <- simState }()
		} else {
			go command.Modifier(&simState, command.Ret)
		}
	}
}

// Updates simulator state based off of incoming json messages
func update_sim_state(msg <-chan []byte, simState chan<- SimStateModCommand) {
	time.Sleep(1 * time.Second)
	for {
		message := <-msg

		var tmpSimState SimulatorState

		err := json.Unmarshal(message, &tmpSimState)
		if err != nil {
			log.Println(err)
		}

		zeroes := 0
		for i := 0; i < len(tmpSimState.State); i++ {
			if tmpSimState.State[i].Count == 0 {
				zeroes++
			}
		}

		log.Printf("tmpSimState length: %d\nzeroes: %d", len(tmpSimState.State), zeroes)

		simState <- SimStateModCommand{
			false,
			func(simState *SimulatorState, ret chan<- SimulatorState) {
				for _, e := range tmpSimState.State {
					(*simState).State[e.TrafficLight] = e
				}
			},
			nil,
		}
	}
}
