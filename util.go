package main

func set_all_red(contrState *ControllerState) {
	for i := 0; i < len((*contrState).State); i++ {
		if (*contrState).State[i].Status == "green" {
			(*contrState).State[i] = ControllerStateSub{i, "orange"}
		} else {
			(*contrState).State[i] = ControllerStateSub{i, "red"}
		}
	}
}