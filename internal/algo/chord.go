package algo

import (
	model "comp-math-2/internal"
	"math"
)

func SolveChord(eq model.NonlinearEquation) (model.Solution, error) {
	f := eq.F
	a := eq.A
	b := eq.B
	eps := eq.Eps
	iterations := 0

	x := a - (b-a)*f(a)/(f(b)-f(a))

	lastX := x

	for ; iterations < 10_000; iterations++ {
		if f(a)*f(b) < 0 {
			b = x
		} else {
			a = x
		}

		x = a - (b-a)*f(a)/(f(b)-f(a))

		if math.Abs(f(x)) <= eps && math.Abs(x-lastX) <= eps {
			break
		}

		lastX = x
	}

	return model.Solution{
		x,
		f(x),
		iterations,
	}, nil
}
