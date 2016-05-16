package gago

import (
	"math/rand"
	"testing"
	"time"
)

var crossovers = []struct {
	crossover Crossover
	init      Initializer
}{
	{CPoint{}, IFUniform{-5.0, 5.0}},
	{CFUniform{}, IFUniform{-5.0, 5.0}},
	{CFProportionate{NbParents: 2}, IFUniform{-5.0, 5.0}},
	{CPMX{}, ISUnique{[]string{"A", "B", "C", "D"}}},
}

func TestCrossovers(t *testing.T) {
	var (
		source    = rand.NewSource(time.Now().UnixNano())
		generator = rand.New(source)
		nbIndis   = 5
		nbGenes   = 4
		nbParents = 2
		selector  = STournament{2}
	)
	for _, c := range crossovers {
		var indis = makeIndividuals(nbIndis, nbGenes)
		var parents = selector.Apply(nbParents, indis, generator)
		// Assign genomes
		for _, indi := range indis {
			c.init.apply(&indi, generator)
		}
		var offsprings = c.crossover.Apply(parents, generator)
		// Check the offspring has a valid genome
		for _, offspring := range offsprings {
			if len(offspring.Genome) != nbGenes {
				t.Error("Crossover operator generated genome of invalid length")
			}
			// Check the offspring doesn't share a pointer with the original population
			for _, indi := range indis {
				if &offspring == &indi {
					t.Error("The offspring shares a pointer with another individual")
				}
			}
		}
	}
}
