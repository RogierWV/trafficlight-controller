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
	// return (ctr_sub.TrafficLight == 0 && ctr_sub.Status == "")
	// var ret []ControllerStateSub
	ret := make([]ControllerStateSub, 29)
	i := 0
	for _,e:=range ctr_arr {
		if e.TrafficLight != 0 && e.Status != "" {
			ret[i] = e
			i++
		}
	}
	return ret
}

func (ctr ControllerState) MarshalJSON() ([]byte, error) {
	// if ctr_sub.TrafficLight != 0 && ctr_sub.Status != "" {
	// 	return json.Marshal(struct{
	// 			TrafficLight int     `json:"trafficLight"`
	// 			Status       *string `json:"status"`
	// 		}{ctr_sub.TrafficLight,&ctr_sub.Status})
	// } else {
	// 	return json.Marshal(struct{d int `json:"-"`}{0})
	// }
	// var ret []byte
	// var err error
	// var b []byte
	// for _,el := range ctr.State {
	// 	if !el.Empty() {
	// 		b,err = json.Marshal(el)
	// 		if err != nil {
	// 			break
	// 		}
	// 		for _,x := range b {
	// 			ret = append(ret[:], x)
	// 		}
	// 	}
	// }
	// return ret, err
	return json.Marshal(struct{State []ControllerStateSub}{Filter(ctr.State)})
}