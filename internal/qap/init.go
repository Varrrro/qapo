package qap

import "math/rand"

// RandomPopulation of permutations with length n.
func RandomPopulation(nPerms, nValues int) []*Permutation {
	p := make([]int, nValues)
	for i := range p {
		p[i] = i
	}

	pop := make([]*Permutation, nPerms)
	for i := range pop {
		pop[i] = NewPermutation(nValues)
		copy(pop[i].Values, p)

		for j := 0; j <= nValues-2; j++ {
			k := j + randomIndex(j, nValues)
			pop[i].Values[j], pop[i].Values[k] = pop[i].Values[k], pop[i].Values[j]
		}
	}

	return pop
}

func randomIndex(begin, end int) int {
	return rand.Intn(end - begin)
}
