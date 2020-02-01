package replacement

import (
	"sort"
	"sync"

	"github.com/varrrro/qapo/internal/qap"
)

// Elitist replacement of a generation.
func Elitist(prev, new []*qap.Permutation, nElite int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// Sort previous generation by fitness
		sort.Slice(prev, func(i, j int) bool {
			return prev[i].Fitness < prev[j].Fitness
		})
		wg.Done()
	}()
	go func() {
		// Sort new generation by fitness
		sort.Slice(new, func(i, j int) bool {
			return new[i].Fitness < new[j].Fitness
		})
		wg.Done()
	}()
	wg.Wait()

	// Replace previous generation except elite
	copy(prev[nElite:], new[:len(prev)-nElite])
}
