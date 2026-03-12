package algo

import (
	"comp-math-2/internal/numeric"
	"fmt"
	"math"
)

func SolveSecant(eq numeric.NonlinearEquation) (numeric.Solution, error) {
	f := eq.F
	a := eq.A
	b := eq.B
	eps := eq.Eps

	prevX := (b + a) / 2.0

	x := prevX + eps

	for i := 1; i <= 10_000; i++ {
		temp := x
		x = x - (x-prevX)/(f(x)-f(prevX))*f(x)

		if math.Abs(x-prevX) <= eps {
			return numeric.Solution{
				X:          x,
				Y:          f(x),
				Iterations: i,
			}, nil
		}

		prevX = temp
	}

	return numeric.Solution{}, fmt.Errorf("maximum number of iterations reached")
}
