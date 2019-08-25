package webassign

import "fmt"

func Title(title string) {
	fmt.Printf("\n### %v ###################\n\n", title)
}
func Solve(problem string, solution interface{}) {
	fmt.Printf("%v) %v\n", problem, solution)
}
