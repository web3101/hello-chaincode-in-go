package main

import (
    "fmt"
    "github.com/cdipaolo/goml/base"
    "github.com/cdipaolo/goml/linear"
)

func main() {
    X := base.F64MatrixFromData(
        []float64{1, 2, 3, 4, 5},
        []float64{1, 2, 3, 4, 5},
    )
    y := base.F64VectorFromData([]float64{2, 4, 6, 8, 10})

    model := linear.NewLinearRegression()
    model.Fit(X, y)

    yPred := model.Predict(X)

    coefficients := model.Coefficients
    intercept := model.Intercept

    fmt.Println("Coefficients:", coefficients)
    fmt.Println("Intercept:", intercept)
    fmt.Println("Predicted Values:", yPred)
}
