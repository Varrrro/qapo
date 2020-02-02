package heuristics

import (
	"sync"

	"github.com/varrrro/qapo/internal/qap"
)

// HillClimbing heuristic algorithm.
func HillClimbing(perms []*qap.Permutation, fitnessDiff func(*qap.Permutation, int, int) int) {
	var wg sync.WaitGroup
	for _, p := range perms {
		wg.Add(1)
		go func(p *qap.Permutation) {
			// Create copy of the base permutation
			tmp := qap.NewPermutation(len(p.Values))
			copy(tmp.Values, p.Values)

			// Apply hill climbing algorithm
			for i := range p.Values {
				for j := range p.Values[i+1:] {
					diff := fitnessDiff(p, i, j)
					if diff < 0 {
						p.Values[i], p.Values[j] = p.Values[j], p.Values[i]
						p.Fitness += diff
						wg.Done()
						return
					}
				}
			}
			wg.Done()
		}(p)
	}
	wg.Wait()
}
