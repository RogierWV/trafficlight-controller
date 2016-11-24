package main

type ControllerState struct {
	State []ControllerStateSub `json:"state"`
}

type ControllerStateSub struct {
	TrafficLight int    `json:"trafficLight"`
	Status       string `json:"status"`
}

type SimulatorState struct {
	State []SimulatorStateSub `json:"state"`
}

type SimulatorStateSub struct {
	TrafficLight int `json:"trafficLight"`
	Count        int `json:"count"`
}

type ContrStateModCommand struct {
	ReadOnly bool
	Modifier func(contrState *ControllerState, ret chan<- ControllerState)
	Ret      chan<- ControllerState
}

type SimStateModCommand struct {
	ReadOnly bool
	Modifier func(simState *SimulatorState, ret chan<- SimulatorState)
	Ret      chan<- SimulatorState
}
