package main

func manage_state(queue <-chan StateModCommand) {
	contrState := ControllerState{ State: make ([]ControllerStateSub, 40) }
	set_all_red(&contrState)
	for {
		select {
			case command := <- queue:
				if command.ReadOnly {
					command.Ret<-contrState
				} else {
					command.Modifier(&contrState, command.Ret)
				}
		}
	}
}

func set_all_red(contrState *ControllerState) {
	for i := 0; i < len((*contrState).State); i++ {
		if (*contrState).State[i].Status == "green" {
			(*contrState).State[i].Status = "orange"
		} else {
			(*contrState).State[i].Status = "red"
		}
	}
}