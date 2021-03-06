package qap

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadData from file.
func ReadData(path string) (n int, w, d [][]int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, nil, err
	}

	w = make([][]int, n)
	for i := range w {
		w[i] = make([]int, n)
		for j := range w[i] {
			scanner.Scan()
			v, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, nil, nil, err
			}
			w[i][j] = v
		}
	}

	d = make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			scanner.Scan()
			v, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, nil, nil, err
			}
			d[i][j] = v
		}
	}

	return n, w, d, nil
}

// WritePermutation values to file.
func WritePermutation(path string, perm *Permutation) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, p := range perm.Values {
		_, err := file.WriteString(fmt.Sprintf("%d ", p))
		if err != nil {
			return err
		}
	}

	return nil
}
