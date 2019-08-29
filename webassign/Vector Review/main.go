package main

import (
	"github.com/manazari/physics"
	wa "github.com/manazari/physics/webassign"
)

var solutions = wa.NewSolutionSheet("answers.txt")

func main() {
	solutions.Title("VECTOR REVIEW")

	section6()
	section7()
	section8()
}

func section6() {
	var (
		a = physics.Vector(3.62, 1.70)
		b = physics.Vector(1.62, 3.50)
	)

	c := b.Subtr(a)
	solutions.Append(
		"\n6a)\t", c,
		"\n6b)\t", "Mag:", c.Mag(), "  Dir:", c.Theta(),
		"\n6c)\t", c.Unit(),
	)
}

func section7() {
	var (
		a = physics.VectorMagDir(10.8, 30)
		b = physics.VectorMagDir(5.2, 53)
		c = physics.VectorMagDir(12.2, 360-60)
		d = physics.VectorMagDir(20.7, 180-37)
		f = physics.VectorMagDir(20.7, 180+30)

		i = physics.VectorMagDir(1, 0)
		j = physics.VectorMagDir(1, 90)
	)
	
	solutions.Append(
		"\n7a)\t", a.Dot(c),
		"\n7a)\t", a.Dot(c),
		"\n7b)\t", a.Dot(f),
		"\n7c)\t", d.Dot(c),
		"\n7d)\t", a.Dot(f.Add(c.Scale(2))),
		"\n7e)\t", i.Dot(b),
		"\n7f)\t", j.Dot(b),
		"\n7g)\t", i.Scale(3).Subtr(j).Dot(b),
		"\n7h)\t", b.Unit().Dot(b),
	)
}

func section8() {
	var (
		a = physics.VectorMagDir(10.1, 30)
		b = physics.VectorMagDir(5.5, 53)
		c = physics.VectorMagDir(12.4, 360-60)
		d = physics.VectorMagDir(20.1, 180-37)
		f = physics.VectorMagDir(20.1, 180+30)

		i = physics.VectorMagDir(1, 0)
		j = physics.VectorMagDir(1, 90)
	)

	solutions.Append(
		"\n8a)\t", a.Cross(c),
		"\n8a)\t", a.Cross(c),
		"\n8b)\t", a.Cross(f),
		"\n8c)\t", d.Cross(c),
		"\n8d)\t", a.Cross(f.Add(c.Scale(2))),
		"\n8e)\t", i.Cross(b),
		"\n8f)\t", j.Cross(b),
		"\n8g)\t", i.Scale(3).Subtr(j).Cross(b),
		"\n8h)\t", b.Unit().Cross(b),
	)
}