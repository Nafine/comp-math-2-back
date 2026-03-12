package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"fmt"
	"math"
)

func SolveSystem(eq numeric.NonlinearSystem) (numeric.SystemSolution, error) {
	coords := numeric.Coordinates{X: eq.StartCoordinates.X, Y: eq.StartCoordinates.Y}
	var prevCoords numeric.Coordinates
	maxIter := 1000

	for iter := 1; iter <= maxIter; iter++ {
		prevCoords = coords
		f1 := eq.F1(coords)
		f2 := eq.F2(coords)

		J11 := derivate.DerivXAt(eq.F1, coords)
		J12 := derivate.DerivYAt(eq.F1, coords)
		J21 := derivate.DerivXAt(eq.F2, coords)
		J22 := derivate.DerivYAt(eq.F2, coords)

		det := J11*J22 - J12*J21

		if math.Abs(det) < 1e-12 {
			return numeric.SystemSolution{}, fmt.Errorf("jacobian is singular at point (%f, %f)", coords.X, coords.Y)
		}

		deltaX := -(f1*J22 - f2*J12) / det
		deltaY := -(J11*f2 - J21*f1) / det

		coords.X += deltaX
		coords.Y += deltaY

		if math.IsNaN(coords.X) || math.IsInf(coords.X, 0) {
			return numeric.SystemSolution{}, fmt.Errorf("method diverged (NaN/Inf encountered)")
		}

		// Проверка на сходимость
		tolerX := math.Abs(coords.X - prevCoords.X)
		tolerY := math.Abs(coords.Y - prevCoords.Y)

		if tolerX <= eq.Eps && tolerY <= eq.Eps {
			return numeric.SystemSolution{
				X:          coords.X,
				Y:          coords.Y,
				Iterations: iter,
				Dx:         tolerX,
				Dy:         tolerY,
			}, nil
		}
	}

	return numeric.SystemSolution{}, fmt.Errorf("maximum number of iterations reached")
}
