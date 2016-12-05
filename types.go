package main

import "encoding/json"

type ControllerState struct {
	State []ControllerStateSub `json:"state,omitempty"`
}

type ControllerStateSub struct {
	TrafficLight int    `json:"trafficLight,omitempty"`
	Status       string `json:"status,omitempty"`
}

type SimulatorState struct {
	State []SimulatorStateSub `json:"state,omitempty"`
}

type SimulatorStateSub struct {
	TrafficLight int `json:"trafficLight,omitempty"`
	Count        int `json:"count,omitempty"`
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

type WL struct {
	ID     int
	Weight int
	Time   int
}

func Filter(ctr_arr []ControllerStateSub) []ControllerStateSub {
	ret := make([]ControllerStateSub, 29)
	i := 0
	for _, e := range ctr_arr {
		if e.TrafficLight != 0 && e.Status != "" {
			ret[i] = e
			i++
		}
	}
	return ret
}

func (ctr ControllerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		State []ControllerStateSub `json:"state"`
	}{Filter(ctr.State)})
}
