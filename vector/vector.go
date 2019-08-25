package vector

import (
	"fmt"
	"math"
)

func sin(theta float64) float64   { return math.Sin(2 * math.Pi * theta / 360) }
func cos(theta float64) float64   { return math.Cos(2 * math.Pi * theta / 360) }
func atan(dy, dx float64) float64 { return math.Atan2(dy, dx) * 360 / (2 * math.Pi) }
func clamp(x float64) float64     { return math.Round(x*1e5) / 1e5 }

type Vect struct {
	Mag, Angle float64
}

type VectIJ struct {
	I, J float64
}

func (v Vect) VectIJ() VectIJ {
	i, j := v.Mag * cos(v.Angle), v.Mag * sin(v.Angle)
	return VectIJ{i ,j}
}

func (v VectIJ) Vect() Vect {
	mag, angle := math.Hypot(v.I, v.J), atan(v.J, v.I)
	return Vect{mag, angle}
}

func (v Vect) String() string {
	return fmt.Sprintf("mag: %2.2f, dir: %5.2f", v.Mag, v.Angle)
}

func (v VectIJ) String() string {
	return fmt.Sprintf("%4.2f\\ivec + %4.2f\\jvec", v.I, v.J)
}

func Cross(a, b Vect) string {
	theta := b.Angle - a.Angle
	cross := a.Mag * b.Mag * sin(theta)
	return fmt.Sprintf("%8.2f\\kvec", cross)
}

func Dot(a, b Vect) string {
	theta := b.Angle - a.Angle
	dot := a.Mag * b.Mag * cos(theta)
	return fmt.Sprintf("%8.2f", dot)
}

func (v Vect) Scale(x float64) Vect {
	a := Vect{v.Mag, v.Angle}
	if x < 0 {
		a.Angle = math.Mod(a.Angle+180, 360)
		x *= -1
	}
	a.Mag *= x
	return a
}

func (v Vect) Unit() Vect {
	return Vect{1, v.Angle}
}

func Add(a, b Vect) Vect {
	ai, aj := a.VectIJ().I, a.VectIJ().J
	bi, bj := b.VectIJ().I, b.VectIJ().J

	ci, cj := ai+bi, aj+bj
	cMag := clamp(math.Hypot(ci, cj))
	cAngle := clamp(atan(cj, ci))

	return Vect{cMag, cAngle}
}

func Subtract(a, b Vect) Vect { return Add(a, b.Scale(-1)) }