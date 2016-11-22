package main

import (
	"time"
)

func _timer(grID int, out chan<- bool, state chan<- StateModCommand, colour string) {
	time.Sleep(3*time.Second)
	state <- StateModCommand {
		false,
		func(contrState *ControllerState, ret chan<- ControllerState){
			for i := 0; i < len(lightGroups[grID]); i++ {
				(*contrState).State[lightGroups[grID][i]].Status = colour
			}
		},
		nil,
	}
	out <- true
}

func timer (grID int, out chan<- bool, state chan<- StateModCommand) {
	_timer(grID, out, state, "yellow")
	_timer(grID, out, state, "red")
}