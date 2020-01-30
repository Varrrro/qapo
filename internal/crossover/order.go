package crossover

import (
	"github.com/varrrro/qapo/internal/qap"
	"math/rand"
	"sync"
)

// Order crossover of the given parents.
func Order(perms []*qap.Permutation) {
	var wg sync.WaitGroup
	for i := 0; i < len(perms); i += 2 {
		wg.Add(1)
		go func(i1, i2 int) {
			p1 := perms[i1]
			p2 := perms[i2]

			size := len(p1.Values)

			c1 := qap.NewPermutation(size)
			c2 := qap.NewPermutation(size)

			start, end := getIndexes(size)

			copy(c1.Values[start:end], p1.Values[start:end])
			copy(c2.Values[start:end], p2.Values[start:end])

			for pi, c1i, c2i := 0, 0, 0; pi < size; pi++ {
				pIndex := (end + pi) % size
				c1Index := (end + c1i) % size
				c2Index := (end + c2i) % size

				if !isPresent(p2.Values[pIndex], c1.Values[start:end]) {
					c1.Values[c1Index] = p2.Values[pIndex]
					c1i++
				}

				if !isPresent(p1.Values[pIndex], c2.Values[start:end]) {
					c2.Values[c2Index] = p1.Values[pIndex]
					c2i++
				}
			}

			perms[i1] = c1
			perms[i2] = c2
			wg.Done()
		}(i, i+1)
	}
	wg.Wait()
}

func getIndexes(size int) (start, end int) {
	n1 := rand.Intn(size - 1)
	n2 := rand.Intn(size)

	if n1 < n2 {
		return n1, n2
	}
	return n2, n1
}

func isPresent(x int, slice []int) bool {
	for _, v := range slice {
		if v == x {
			return true
		}
	}
	return false
}
