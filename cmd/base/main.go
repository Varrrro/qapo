package main

import (
	"math"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/varrrro/qapo/internal/crossover"
	"github.com/varrrro/qapo/internal/mutation"
	"github.com/varrrro/qapo/internal/qap"
	"github.com/varrrro/qapo/internal/replacement"
	"github.com/varrrro/qapo/internal/selection"
)

const (
	dataPath = "data/tai256c.dat"
	nGens    = 1000
)

func init() {
	// Set log formatter
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	// Write logs to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	// Read data from file
	n, w, d, err := qap.ReadData(dataPath)
	if err != nil {
		log.WithField("path", dataPath).WithError(err).Fatal("Can't read data file")
	}

	start := time.Now()
	// Generate initial population and calculate its fitness
	pop := qap.RandomPopulation(100, n)
	qap.CalculateFitness(pop, w, d)
	// Start generations
	for i := 0; i < nGens; i++ {
		// Select parents
		tmp := selection.Tournament(pop, 10, 100)
		// Apply crossover to parents
		crossover.Order(tmp)
		// Mutate children
		mutation.RandomSwap(tmp)
		// Calculate children fitness
		qap.CalculateFitness(tmp, w, d)
		// Replace population
		replacement.Elitist(pop, tmp, 10)

		log.WithFields(log.Fields{
			"gen":  i,
			"best": getBest(pop).Fitness,
		}).Info("Finished generation")
	}

	elapsed := time.Since(start)
	log.WithField("seconds", elapsed.Seconds()).Info("Finished execution")

	qap.WritePermutation("results/base", getBest(pop))
}

func getBest(perms []*qap.Permutation) *qap.Permutation {
	f := math.MaxInt64
	i := -1
	for j, p := range perms {
		if p.Fitness < f {
			f = p.Fitness
			i = j
		}
	}
	return perms[i]
}
