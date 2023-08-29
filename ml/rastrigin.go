package main

import (
    "fmt"
    m "math"
    "math/rand"

    "github.com/MaxHalford/eaopt"
)

func Rastrigin(x []float64) (y float64) {
    y = 10 * float64(len(x))
    for _, xi := range x {
        y += m.Pow(xi, 2) - 10*m.Cos(2*m.Pi*xi)
    }
    return y
}

func main() {
    // Instantiate DiffEvo
    var oes, err = eaopt.NewDefaultOES()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Fix random number generation
    oes.GA.RNG = rand.New(rand.NewSource(42))

    // Run minimization
    _, y, err := oes.Minimize(Rastrigin, 2, []float64{0.1})
    if err != nil {
        fmt.Println(err)
        return
    }

    // Output best encountered solution
    fmt.Printf("Found minimum of %.5f, the global minimum is 0\n", y)
}