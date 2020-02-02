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

// CalculateDifference in fitness swapping two positions in a permutation.
func CalculateDifference(perm *Permutation, i, j int, w, d [][]int) int {
	x := perm.Values[i]
	y := perm.Values[j]
	sum := 0
	for k, z := range perm.Values {
		if k != i && k != j {
			sum += w[k][i]*(d[z][y]-d[z][x]) +
				w[k][j]*(d[z][x]-d[z][y]) +
				w[i][k]*(d[y][z]-d[x][z]) +
				w[j][k]*(d[x][z]-d[y][z])
		}
	}
	return sum
}
