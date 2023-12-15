package main

import (
	"math/rand"
	"time"
)

func process_simstate(out chan<- bool, contrState chan<- ContrStateModCommand, simState chan<- SimStateModCommand) {
	time.Sleep(1 * time.Second)
	simStateRet := make(chan SimulatorState)
	for {
		// log.Println("process iteration")
		simState <- SimStateModCommand{true, nil, simStateRet} // request current simstate

		highestTotal := 0
		groupId := -1

		tmpLights := make([][]WL, len(newLightGroups)) // build up arrays for current iteration
		if *randomise {                                // -r flag enabled
			perm := rand.Perm(len(newLightGroups)) // randomize the lights
			for i, v := range perm {
				tmpLights[i] = newLightGroups[v]
			}
		} else {
			tmpLights = newLightGroups // set to the values from lights.go
		}

		tmpSimState := <-simStateRet // actually fetch current simstate

		for i := 0; i < len(tmpLights); i++ {
			total := 0
			for j := 0; j < len(tmpLights[i]); j++ {
				count := tmpSimState.State[tmpLights[i][j].ID].Count
				weight := tmpLights[i][j].Weight
				time := &tmpLights[i][j].Time
				(*time)++
				// log.Printf("{\"count\": %d, \"weight\": %d, \"time\": %d}", count, weight, *time)
				total += count * weight * (*time) // compute priority
			}
			if total > highestTotal { // this iteration has the highest priority so far
				highestTotal = total
				groupId = i
			}
		}

		// log.Printf("highestTotal = %d\ngroupId = %d", highestTotal, groupId)

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
