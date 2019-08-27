package physics

import "math"

func toDeg(x float64) (deg float64) {
	deg = x * 180 / math.Pi
	return
}

func toRad(x float64) (rad float64) {
	rad = x * math.Pi / 180
	return
}

func contain360(x float64) (deg float64) {
	deg = math.Mod(x, math.Mod(x, 360)+360)
	return
}

func sin(x float64) float64 {
	return math.Sin(toRad(x))
}

func cos(x float64) float64 {
	return math.Cos(toRad(x))
}
