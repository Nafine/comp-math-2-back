package main

import (
	model "comp-math-2/internal"
	"comp-math-2/internal/algo"
	"fmt"
)

func main() {

	eq := model.NonlinearEquation{
		F: func(x float64) float64 {
			return x*x*x + 2.84*x*x - 5.606*x - 14.766
		},
		A:   -4.0,
		B:   -3.0,
		Eps: 0.01,
	}

	fmt.Println(algo.SolveChord(eq))
	fmt.Println(algo.SolveSimpleIteration(eq))
	fmt.Println(algo.SolveSplitting(eq))

	eq.A = -3
	eq.B = 0

	fmt.Println(algo.SolveNewton(eq))
}
