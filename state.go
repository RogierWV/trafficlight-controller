package main

import (
	"log"
	"encoding/json"
)

func manage_controller_state(queue <-chan ContrStateModCommand) {
	contrState := ControllerState{ State: make ([]ControllerStateSub, 50) }
	set_all_red(&contrState)
	for {
		command := <- queue
		// log.Println(command)
		if command.ReadOnly {
			go func() {command.Ret<-contrState}()
		} else {
			go command.Modifier(&contrState, command.Ret)
		}
	}
}

func set_all_red(contrState *ControllerState) {
	for i := 0; i < len((*contrState).State); i++ {
		if (*contrState).State[i].Status == "green" {
			(*contrState).State[i] = ControllerStateSub{i, "orange"}
		} else {
			(*contrState).State[i] = ControllerStateSub{i, "red"}
		}
	}
}

func manage_sim_state(queue <-chan SimStateModCommand) {
	simState := SimulatorState { State: make([]SimulatorStateSub, 50) }
	for i := 0; i < len(simState.State); i++ {
		simState.State[i] = SimulatorStateSub{i,0}
	}

	for {
		command := <- queue
		if command.ReadOnly {
			go func() {command.Ret <- simState}()
		} else {
			go command.Modifier(&simState, command.Ret)
		}
	}
}

func update_sim_state(msg <-chan []byte, simState chan<- SimStateModCommand) {
	for {
		message := <-msg

		var tmpSimState SimulatorState

		err := json.Unmarshal(message, &tmpSimState)
		if err != nil {
			log.Println(err)
		}

		simState <- SimStateModCommand {
			false,
			func(simState *SimulatorState, ret chan<- SimulatorState) {
				for _,e := range tmpSimState.State {
					(*simState).State[e.TrafficLight] = e
				}
			},
			nil,
		}
	}
}