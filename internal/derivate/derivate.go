package derivate

func DerivAt(f func(float64) float64, x float64) float64 {
	h := 1e-5
	return (f(x+h) - f(x-h)) / (2 * h)
}

func Derivate(f func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return DerivAt(f, x)
	}
}
