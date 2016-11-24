package main

import (
	"time"
)

func _timer(grID int, out chan<- bool, state chan<- ContrStateModCommand, colour string) {
	// if grID == -1 {
	// 	return
	// }
	time.Sleep(3 * time.Second)
	state <- ContrStateModCommand{
		false,
		func(contrState *ControllerState, ret chan<- ControllerState) {
			for i := 0; i < len(lightGroups[grID]); i++ {
				(*contrState).State[lightGroups[grID][i]].Status = colour
			}
		},
		nil,
	}
	out <- true
}

func timer(grID int, out chan<- bool, state chan<- ContrStateModCommand) {
	_timer(grID, out, state, "yellow")
	_timer(grID, out, state, "red")
	time.Sleep(time.Duration(*redTime) * time.Second)
}
