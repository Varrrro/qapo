package heuristics

import (
	"sync"

	"github.com/varrrro/qapo/internal/qap"
)

// HillClimbing heuristic algorithm.
func HillClimbing(perms []*qap.Permutation, fitness func(*qap.Permutation)) {
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
					tmp.Values[i], tmp.Values[j] = tmp.Values[j], tmp.Values[i]
					fitness(tmp)

					if tmp.Fitness < p.Fitness {
						copy(p.Values, tmp.Values)
						p.Fitness = tmp.Fitness

						wg.Done()
						return
					}

					tmp.Values[i], tmp.Values[j] = tmp.Values[j], tmp.Values[i]
				}
			}
			wg.Done()
		}(p)
	}
	wg.Wait()
}
