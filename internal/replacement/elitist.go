package replacement

import (
	"github.com/varrrro/qapo/qap"
	"sort"
	"sync"
)

// Elitist replacement of a generation.
func Elitist(prev, new []*qap.Permutation, nElite int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		sort.Slice(prev, func(i, j int) bool {
			return prev[i].Fitness < prev[j].Fitness
		})
		wg.Done()
	}()
	go func() {
		sort.Slice(new, func(i, j int) bool {
			return new[i].Fitness < new[j].Fitness
		})
		wg.Done()
	}()
	wg.Wait()

	prev[nElite:] = new[:len(prev)-nElite]
}
