package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

// A Coord2D is a coordinate in two dimensions.
type Coord2D struct {
	X float64
	Y float64
}

// Evaluate evalutes a Bohachevsky function at the current coordinates.
func (c *Coord2D) Evaluate() (float64, error) {
	z := c.X*c.X + 2*c.Y*c.Y - 0.3*math.Cos(3*math.Pi*c.X) - 0.4*math.Cos(4*math.Pi*c.Y) + 0.7
	return z, nil
}

// Mutate replaces one of the current coordinates with a random value in [-100, -100].
func (c *Coord2D) Mutate(rng *rand.Rand) {
	if rng.Intn(2) == 0 {
		c.X = rng.Float64()*200.0 - 100.0
	} else {
		c.Y = rng.Float64()*200.0 - 100.0
	}
}

// Crossover does nothing.  It is defined only so *Coord2D implements the eaopt.Genome interface.
func (c *Coord2D) Crossover(other eaopt.Genome, rng *rand.Rand) {}

// Clone returns a copy of a *Coord2D.
func (c *Coord2D) Clone() eaopt.Genome {
	return &Coord2D{X: c.X, Y: c.Y}
}

func main() {
	// Hill climbing is implemented as a GA using the ModMutationOnly model
	// with the Strict option.
	cfg := eaopt.NewDefaultGAConfig()
	cfg.Model = eaopt.ModMutationOnly{Strict: true}
	cfg.NGenerations = 9999

	// Add a custom callback function to track progress.
	minFit := math.MaxFloat64
	cfg.Callback = func(ga *eaopt.GA) {
		hof := ga.HallOfFame[0]
		fit := hof.Fitness
		if fit == minFit {
			// Output only when we make an improvement.
			return
		}
		best := hof.Genome.(*Coord2D)
		fmt.Printf("Best fitness at generation %4d: %10.5f at (%9.5f, %9.5f)\n",
			ga.Generations, fit, best.X, best.Y)
		minFit = fit
	}

	// Run the hill-climbing algorithm.
	ga, err := cfg.NewGA()
	if err != nil {
		panic(err)
	}
	err = ga.Minimize(func(rng *rand.Rand) eaopt.Genome {
		return &Coord2D{
			X: rng.Float64()*200.0 - 100.0,
			Y: rng.Float64()*200.0 - 100.0,
		}
	})
	if err != nil {
		panic(err)
	}

	// Output the best encountered solution.
	best := ga.HallOfFame[0].Genome.(*Coord2D)
	fmt.Printf("Found a minimum at (%.5f, %.5f).\n", best.X, best.Y)
	fmt.Println("The global minimum is known to lie at (0, 0).")
}