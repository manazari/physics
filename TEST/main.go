package main

import (
	"fmt"

	p "github.com/manazari/physics"
)

func main() {
	a, b := p.Vector(1, 9), p.Vector(5, 3)
	fmt.Println(a, b)
	fmt.Printf("Dimensions: %v\n", a.Dimensions())
	fmt.Printf("Unit: %v | Unit Mag: %v\n", a.Unit(), a.Unit().Mag())
	fmt.Printf("Theta: %v\n", a.Theta())
	fmt.Printf("Mag: %v\n", a.Mag())
	// for i := 0.0; i < 1; i += .01 {
	// 	a.SetMag(i)
	// 	fmt.Printf("Mag: %v\n", a.Mag())
	// }
	a.SetMag(29)
	fmt.Printf("Mag: %v\n", a.Mag())

	fmt.Println(a.Scale(10))
	a.Set(-90, 90, 90)
	b.SetMagDir(0.8391, 45)
	b = append(b, -1)
	fmt.Println(a, b, a.AngleWith(b))
	// a.SetTheta(180)
	// fmt.Println(a)
	// fmt.Printf("Theta: %v\n", a.Theta())

	// fmt.Println(b.Theta(), b.Mag())
	// fmt.Println(a.Dot(b))

	c := p.Vector(1.5381, -0.0001, -3.9)
	d := p.Vector(0.98, -2.51, -45.2)
	fmt.Printf("%v\n%v", c.WAFormat(), d.WAFormat())
}
