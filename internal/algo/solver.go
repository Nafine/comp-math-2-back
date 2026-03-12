package algo

import (
	"comp-math-2/internal/numeric"
	"errors"
)

var methods = map[string]func(equation numeric.NonlinearEquation) (numeric.Solution, error){
	"chord":     SolveChord,
	"secant":    SolveSecant,
	"iteration": SolveSimpleIteration,
}

func SolveSingle(method string, eq numeric.NonlinearEquation) (numeric.Solution, error) {
	if !rootExists(eq) {
		return numeric.Solution{}, errors.New("no roots exists on the given interval")
	}

	if eq.A >= eq.B {
		return numeric.Solution{}, errors.New("a must be higher than b")
	}

	if eq.Eps <= 0 {
		return numeric.Solution{}, errors.New("eps must be greater than zero")
	}

	return methods[method](eq)
}

func rootExists(eq numeric.NonlinearEquation) bool {
	return eq.F(eq.A)*eq.F(eq.B) < 0
}
