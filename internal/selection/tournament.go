package selection

import (
	"math/rand"
	"sort"
	"sync"

	"github.com/varrrro/qapo/internal/qap"
)

// Tournament selection with one winner per tournament.
// k is the number of candidates.
// n is the number of total winners.
func Tournament(perms []*qap.Permutation, k, n int) []*qap.Permutation {
	var wg sync.WaitGroup
	sel := make([]*qap.Permutation, n)
	for i := range sel {
		wg.Add(1)
		go func(i int) {
			// Create slice with random participants
			tmp := make([]*qap.Permutation, k)
			for j := range tmp {
				tmp[j] = perms[rand.Intn(len(perms))]
			}

			// Sort slice by fitness
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i].Fitness < tmp[j].Fitness
			})

			// Lowest fitness is the winner
			sel[i] = tmp[0]
			wg.Done()
		}(i)
	}
	wg.Wait()

	return sel
}
