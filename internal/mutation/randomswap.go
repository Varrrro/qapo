package mutation

import (
	"math/rand"
	"sync"

	"github.com/varrrro/qapo/internal/qap"
)

// RandomSwap mutation of the given children.
func RandomSwap(perms []*qap.Permutation) {
	var wg sync.WaitGroup
	for _, p := range perms {
		// Check if the individual gets mutated
		if allowMutation() == true {
			wg.Add(1)
			go func(p *qap.Permutation) {
				// Select two random indexes
				i1 := rand.Intn(len(p.Values))
				i2 := rand.Intn(len(p.Values))

				// Swap values
				p.Values[i1], p.Values[i2] = p.Values[i2], p.Values[i1]
				wg.Done()
			}(p)
		}
	}
	wg.Wait()
}

func allowMutation() bool {
	if rand.Float64() < 0.8 {
		return true
	}
	return false
}
