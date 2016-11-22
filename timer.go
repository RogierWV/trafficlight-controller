package main

import (
	"time"
)

// func time_to_yellow(id <-chan int, out chan<- bool, state chan<- StateModCommand, red chan<- int) {
// 	for {
// 		tlid := <- id
// 		time.Sleep(3*time.Second)
// 		red<-tlid
// 		state <- StateModCommand {
// 			false,
// 			func(contrState *ControllerState, ret chan<- ControllerState){(*contrState).State[tlid].Status = "yellow"},
// 			nil,
// 		}
// 		out <- true
// 	}
// }

// func time_to_red(id <-chan int, out chan<- bool, state chan<- StateModCommand) {
// 	for {
// 		tlid := <- id
// 		time.Sleep(3*time.Second)
// 		state <- StateModCommand {
// 			false,
// 			func(contrState *ControllerState, ret chan<- ControllerState){(*contrState).State[tlid].Status = "red"},
// 			nil,
// 		}
// 		out <- true
// 	}
// }

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