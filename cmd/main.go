package main

import (
	"math"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/varrrro/qapo/internal/crossover"
	"github.com/varrrro/qapo/internal/mutation"
	"github.com/varrrro/qapo/internal/qap"
	"github.com/varrrro/qapo/internal/replacement"
	"github.com/varrrro/qapo/internal/selection"
)

const dataPath = "../data/tai256.dat"

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
	n, w, d, err := qap.ReadData(dataPath)
	if err != nil {
		log.WithField("path", dataPath).WithError(err).Fatal("Can't read data file")
	}

	po := qap.RandomPopulation(100, n)
	qap.CalculateFitness(po, w, d)
	for i := 0; i < 1000; i++ {
		tmp := selection.Tournament(po, 10, 100)
		crossover.Order(tmp)
		mutation.RandomSwap(tmp)
		qap.CalculateFitness(tmp)
		replacement.Elitist(po, tmp, 10)

		log.WithFields(log.Fields{
			"gen":  i,
			"best": getBest(po).Fitness,
		}).Info("Finished generation")
	}

	log.WithFields(log.Fields{
		"best": getBest(po).Fitness,
	}).Info("Finished execution")
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
