package webassign

import (
	"fmt"
	"os"
)

func CreateSolutionSheet() {
	f, err := os.Create("solutions.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()
}

func AppendToSolutions(text string) {
	if _, err := os.Stat("solutions.md"); os.IsNotExist(err) {
		CreateSolutionSheet()
	}
	
	f, err := os.OpenFile("solutions.md", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		fmt.Println(err)
	}
}

func main() {
	AppendToSolutions("dab")
}

