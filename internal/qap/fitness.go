package qap

import (
	"sync"
)

// CalculateFitness of a set of permutations.
// w is the weights matrix.
// d is the distances matrix.
func CalculateFitness(perms []*Permutation, w, d [][]int) {
	var wg sync.WaitGroup
	for _, p := range perms {
		wg.Add(1)
		go func(p *Permutation) {
			sum := 0
			for i, x := range p.Values {
				for j, y := range p.Values {
					sum += w[i][j] * d[x][y]
				}
			}

			p.Fitness = sum
			wg.Done()
		}(p)
	}
	wg.Wait()
}
