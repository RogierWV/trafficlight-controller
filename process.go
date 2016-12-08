package main

import (
	"time"
	"math/rand"
)

func process_simstate(out chan<- bool, contrState chan<- ContrStateModCommand, simState chan<- SimStateModCommand) {
	time.Sleep(1 * time.Second)
	simStateRet := make(chan SimulatorState)
	for {
		simState <- SimStateModCommand{true, nil, simStateRet}

		highestTotal := 0
		groupId := -1

		tmpSimState := <-simStateRet

		tmpLights := make([][]WL, len(newLightGroups))
		if randomise {
			perm := rand.Perm(len(newLightGroups))
			for i,v := range perm {
				tmpLights[i] = newLightGroups[v]
			}
		} else {
			tmpLights = newLightGroups
		}

		for i := 0; i < len(tmpLights); i++ {
			total := 0
			for j := 0; j < len(tmpLights[i]); j++ {
				count := tmpSimState.State[tmpLights[i][j].ID].Count
				weight := tmpLights[i][j].Weight
				time := &tmpLights[i][j].Time
				(*time)++
				total += count * weight * (*time)
			}
			if total > highestTotal {
				highestTotal = total
				groupId = i
			}
		}

		if groupId != -1 {
			contrState <- ContrStateModCommand{
				false,
				func(contrState *ControllerState, ret chan<- ControllerState) {
					for _, e := range newLightGroups[groupId] {
						(*contrState).State[e.ID] = ControllerStateSub{e.ID, "green"}
					}
				},
				nil,
			}
			for _, e := range newLightGroups[groupId] {
				e.Time = 0
			}
			out <- true
			timer(groupId, out, contrState)
		}
	}
}
