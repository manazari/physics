package physics

import "math"

func clamp(x, min, max float64) float64 {
	return math.Min(math.Max(min, x), max)
}

func trunc(x float64, decimalPoint uint) float64 {
	shift := math.Pow10(int(decimalPoint))
	return math.Round(x * shift) / shift
}