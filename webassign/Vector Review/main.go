package main

import (
	"fmt"

	"github.com/manazari/physics/vector"
	"github.com/manazari/physics/webassign"
)

var solve = webassign.Solve

func main() {
	webassign.CreateSolutionSheet()
	webassign.Title("VECTOR REVIEW")

	section6()
	section7()
	section8()
}

func section6() {
	var (
		a = vector.VectIJ{3.62, 1.70}
		b = vector.VectIJ{1.62, 3.50}
	)

	subtract := vector.Subtract
	c := subtract(b.Vect(), a.Vect())
	solve("6a", c.VectIJ())
	solve("6b", c)
	solve("6c", c.Unit().VectIJ())
	fmt.Println("\n")
	webassign.AppendToSolutions("\n")
}

func section7() {
	var (
		a = vector.Vect{10.8, 30}
		b = vector.Vect{5.2, 53}
		c = vector.Vect{12.2, 360 - 60}
		d = vector.Vect{20.7, 180 - 37}
		f = vector.Vect{20.7, 180 + 30}

		i = vector.Vect{1, 0}
		j = vector.Vect{1, 90}
	)

	dot := vector.Dot
	add := vector.Add
	subtract := vector.Subtract
	solve("7a", dot(a, c))
	solve("7b", dot(a, f))
	solve("7c", dot(d, c))
	solve("7d", dot(a, add(f, c.Scale(2))))
	solve("7e", dot(i, b))
	solve("7f", dot(j, b))
	solve("7g", dot(subtract(i.Scale(3), j), b))
	solve("7h", dot(b.Unit(), b))
	fmt.Println("\n")
	webassign.AppendToSolutions("\n")
}

func section8() {
	var (
		a = vector.Vect{10.1, 30}
		b = vector.Vect{5.5, 53}
		c = vector.Vect{12.4, 360 - 60}
		d = vector.Vect{20.1, 180 - 37}
		f = vector.Vect{20.1, 180 + 30}

		i = vector.Vect{1, 0}
		j = vector.Vect{1, 90}
	)

	cross := vector.Cross
	add := vector.Add
	subtract := vector.Subtract
	solve("8a", cross(a, c))
	solve("8b", cross(a, f))
	solve("8c", cross(d, c))
	solve("8d", cross(a, add(f, c.Scale(2))))
	solve("8e", cross(i, b))
	solve("8f", cross(j, b))
	solve("8g", cross(subtract(i.Scale(3), j), b))
	solve("8h", cross(b.Unit(), b))
	fmt.Println("\n")
	webassign.AppendToSolutions("\n")
}