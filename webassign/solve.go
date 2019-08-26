package webassign

import "fmt"

func Title(title string) {
	text := fmt.Sprintf("\n### %v ###################\n\n", title)
	fmt.Print(text)
	AppendToSolutions(text)
}
func Solve(problem string, solution interface{}) {
	text := fmt.Sprintf("%v) %v\n", problem, solution)
	fmt.Print(text)
	AppendToSolutions(text)
}