package main

// import "log"

func manage_state(queue <-chan StateModCommand) {
	contrState := ControllerState{ State: make ([]ControllerStateSub, 40) }
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