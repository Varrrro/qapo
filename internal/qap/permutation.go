package qap

// Permutation of items that represents a possible solution.
type Permutation struct {
	Values  []int
	Fitness int
}

// NewPermutation with the given size.
func NewPermutation(n int) *Permutation {
	return &Permutation{
		Values: make([]int, n),
	}
}
