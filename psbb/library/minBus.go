package library

import (
	"fmt"
	"sort"
)

func MinBus(fams []int) {
	bus := 0
	visit := make([]bool, len(fams))
	sort.Slice(fams, func(i, j int) bool { return fams[j] < fams[i] })

	for i := 0; i < len(fams); i++ {
		if visit[i] {
			continue
		}
		bus++
		visit[i] = true
		for j := i + 1; j < len(fams); j++ {
			if !visit[j] && fams[i]+fams[j] <= 4 {
				visit[j] = true
				break
			}
		}	
	}

	fmt.Println("Minimum bus required is :", bus)
}
