package main

import "time"

/*
dest := make([]int, len(src))
perm := rand.Perm(len(src))
for i, v := range perm {
    dest[v] = src[i]
}
*/

func process_simstate(out chan<- bool, contrState chan<- ContrStateModCommand, simState chan<- SimStateModCommand) {
	time.Sleep(1 * time.Second)
	simStateRet := make(chan SimulatorState)
	for {
		simState <- SimStateModCommand{true, nil, simStateRet}

		highestTotal := 0
		groupId := -1

		tmpSimState := <-simStateRet

		// temp, fix with above function for more fair scheduling
		forward := true
		if forward{
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
			forward = false
		} else {
			for i := len(newLightGroups); i >= 0; i-- {
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
			forward = true
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
