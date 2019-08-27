package physics

import "math"

type vector []float64

func Vector(components ...float64) (v vector) {
	if len(components) > 3 {
		components = components[:3]
	}
	v = append(vector{}, components...)
	return
}

func (v vector) Dimensions() (dim int) {
	dim = len(v)
	return
}

func (v vector) Unit() (unit vector) {
	unit, mag := v, v.Mag()
	for i := range unit {
		unit[i] /= mag
	}
	return
}

func (v vector) Dot(v2 vector) (dot float64) {
	if v.Dimensions() != v2.Dimensions() {
		panic("Mismatching vector dimensions")
	}
	for i := range v {
		dot += v[i] * v2[i]
	}
	return
}

func (v vector) Cross(v2 vector) (cross vector) {
	if v.Dimensions() != v2.Dimensions() {
		panic("Mismatching vector dimensions")
	}
	for len(v) < 3 {
		v = append(v, 0)
		v2 = append(v2, 0)
	}
	cross = Vector(
		v[2]*v2[3]-v[3]*v2[2],
		v[3]*v2[1]-v[1]*v2[3],
		v[1]*v2[2]-v[2]*v2[1],
	)
	return
}

func (v vector) Mag() (mag float64) {
	for _, component := range v {
		mag += math.Pow(component, 2)
	}
	return
}

func (v vector) Theta() (theta float64) {
	if v.Dimensions() != 2 {
		panic("Can only get theta for 2D vector")
	}
	theta = contain360(toDeg(math.Atan2(v[1], v[0])))
	return
}

func (v *vector) SetMag(x float64) {
	unit := v.Unit()
	for i := range *v {
		(*v)[i] = unit[i] * x
	}
}

func (v *vector) SetTheta(x float64) {
	x = contain360(x)
	if v.Dimensions() != 2 {
		panic("Can only set theta for 2D vector")
	}
	(*v)[0], (*v)[1] = x*cos(v.Theta()), x*sin(v.Theta())
}

func (v vector) Scale(x float64) (new vector) {
	v.SetMag(v.Mag() * x)
	new = v
	return
}

func (v vector) AngleWith(v2 vector) (theta float64) {
	theta = math.Acos(v.Dot(v2) / (v.Mag() * v2.Mag()))
	return
}
