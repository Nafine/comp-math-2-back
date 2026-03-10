package main

import (
	model "comp-math-2/internal"
	"comp-math-2/internal/algo"
	"fmt"
	"math"
)

func main() {
	//eq := model.NonlinearEquation{
	//	F: func(x float64) float64 {
	//		return x*x*x + 2.84*x*x - 5.606*x - 14.766
	//	},
	//	A:   -4.0,
	//	B:   -3.0,
	//	Eps: 0.01,
	//}
	//
	//fmt.Println(algo.SolveChord(eq))
	//fmt.Println(algo.SolveSimpleIteration(eq))
	//fmt.Println(algo.SolveSplitting(eq))
	//
	//eq.A = -3
	//eq.B = 0
	//
	//fmt.Println(algo.SolveNewton(eq))

	eqs := model.NonlinearSystem{
		F1: func(coords model.Coordinates) float64 {
			return math.Cos(coords.X-1) + coords.Y - 0.5
		},
		F2: func(coords model.Coordinates) float64 {
			return coords.X - math.Cos(coords.Y) - 3
		},
		X0:  3.3,
		Y0:  1.2,
		Eps: 0.001,
	}

	fmt.Println(algo.SolveSystem(eqs))
}
