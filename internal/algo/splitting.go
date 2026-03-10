package algo

import (
	model "comp-math-2/internal"
	"math"
)

func SolveSplitting(eq model.NonlinearEquation) (model.Solution, error) {
	f := eq.F
	a := eq.A
	b := eq.B
	eps := eq.Eps
	iterations := 0

	prevX := (b + a) / 2.0

	x := prevX + eps

	for {
		iterations++
		temp := x
		x = x - (x-prevX)/(f(x)-f(prevX))*f(x)

		if math.Abs(x-prevX) <= eps {
			break
		}

		prevX = temp
	}

	return model.Solution{
		X:          x,
		Y:          f(x),
		Iterations: iterations,
	}, nil
}
