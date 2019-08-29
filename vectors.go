package physics

import (
	"fmt"
	"math"
)

type vector []float64

func Vector(components ...float64) (v vector) {
	if len(components) > 3 {
		components = components[:3]
	}
	v = append(vector{}, components...)
	return
}

func VectorMagDir(mag, dir float64) (v vector) {
	v = vector{1, 1}
	v.SetMagDir(mag, dir)
	return
}

func (v vector) Copy() (new vector) {
	new = append(vector{}, v...)
	return
}

func (v vector) Dimensions() (dim int) {
	dim = len(v)
	return
}

func (v vector) Unit() (unit vector) {
	unit, mag := v.Copy(), v.Mag()
	for i := range unit {
		unit[i] /= mag
	}
	return
}

func (v vector) Scale(x float64) (new vector) {
	new = v.Copy()
	new.SetMag(new.Mag() * x)
	return
}

// MAG, DIR ###################################

func (v vector) Mag() (mag float64) {
	for _, component := range v {
		mag = math.Hypot(mag, component)
	}
	mag = trunc(mag, 7)
	return
}

func (v vector) Theta() (theta float64) {
	if v.Dimensions() != 2 {
		panic("Can only get theta for 2D vector")
	}
	theta = toDeg(math.Atan2(v[1], v[0]))
	theta = trunc(theta, 7)
	return
}

func (v vector) MagDir() (vector []float64) {
	vector = []float64{v.Mag(), v.Theta()}
	return
}

func (v *vector) Set(components ...float64) {
	*v = append(vector{}, components...)
}

func (v vector) SetMagDir(mag, dir float64) {
	v.SetMag(mag)
	v.SetTheta(dir)
}

func (v vector) SetMag(x float64) {
	unit, x := v.Unit(), math.Max(1e-7, x)
	for i := range v {
		v[i] = unit[i] * x
	}
}

func (v vector) SetTheta(x float64) {
	x = contain360(x)
	if v.Dimensions() != 2 {
		panic("Can only set theta for 2D vector")
	}
	v[0], v[1] = v.Mag()*cos(x), v.Mag()*sin(x)
	v[0], v[1] = trunc(v[0], 10), trunc(v[1], 10)
}

func (v vector) AngleWith(v2 vector) (theta float64) {
	theta = math.Acos(v.Dot(v2) / (v.Mag() * v2.Mag()))
	theta = toDeg(theta)
	theta = trunc(theta, 7)
	return
}

// Mathematics ###############################################

func (v vector) AssertSameDimensions(v2 vector) {
	if v.Dimensions() != v2.Dimensions() {
		panic("Mismatching vector dimensions")
	}
}

func (v vector) Add(v2 vector) (resultant vector) {
	v.AssertSameDimensions(v2)
	for i := range v {
		resultant = append(resultant, v[i]+v2[i])
	}
	return
}

func (v vector) Subtr(v2 vector) (resultant vector) {
	return v.Add(v2.Scale(-1))
}

func (v vector) Dot(v2 vector) (dot float64) {
	v.AssertSameDimensions(v2)
	for i := range v {
		dot += v[i] * v2[i]
	}
	return
}

func (v vector) Cross(v2 vector) (cross vector) {
	a, b := v.Copy(), v2.Copy()
	if a.Dimensions() != b.Dimensions() {
		panic("Mismatching vector dimensions")
	}

	for len(a) < 3 {
		a = append(a, 0)
		b = append(b, 0)
	}

	cross = Vector(
		a[1]*b[2]-a[2]*b[1],
		a[2]*b[0]-a[0]*b[2],
		a[0]*b[1]-a[1]*b[0],
	)
	return
}

// FORMAT ###############################

func (v vector) String() (str string) {
	suffixes := []string{
		"\\ivec",
		"\\jvec",
		"\\kvec",
	}
	for i := range v {
		if trunc(v[i], 2) == 0 && i+1 != len(v) {
			continue
		}

		str = fmt.Sprintf("%v%+7.2f%v ", str, v[i], suffixes[i])
	}
	return
}
