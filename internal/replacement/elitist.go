package replacement

import (
	"github.com/varrrro/qapo/internal/qap"
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

	copy(prev[nElite:], new[:len(prev)-nElite])
}
