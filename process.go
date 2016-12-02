package main

// import (
// 	"encoding/json"
// 	"log"
// )

func process_simstate(out chan<- bool, contrState chan<- ContrStateModCommand, simState chan<- SimStateModCommand) {
	simStateRet := make(chan SimulatorState)
	for {
		simState <- SimStateModCommand{true, nil, simStateRet}

		highestTotal := 0
		groupId := -1

		tmpSimState := <-simStateRet

		// for i := 0; i < len(lightGroups); i++ {
		// 	total := 0
		// 	for j := 0; j < len(lightGroups[i]); j++ {
		// 		total += tmpSimState.State[lightGroups[i][j]].Count
		// 	}
		// 	if total > highestTotal {
		// 		highestTotal = total
		// 		groupId = i
		// 	}
		// }

		for i := 0; i < len(newLightGroups); i++ {
			total := 0
			for j := 0; j < len(newLightGroups[i]); j++ {
				count := tmpSimState.State[newLightGroups[i][j].ID].Count
				weight := newLightGroups[i][j].Weight
				time := &newLightGroups[i][j].Time
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
