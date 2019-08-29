package webassign

import (
	"fmt"
	"os"
)

var format = fmt.Sprintf

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type solutionSheet struct {
	Name        string
	Body        string
	initialized bool
}

func NewSolutionSheet(name string) solutionSheet {
	s := solutionSheet{name, "", false}
	s.Init()
	return s
}

func (s *solutionSheet) Init() {
	f, err := os.Create(s.Name)
	check(err)
	s.initialized = true
	check(f.Close())
}

func (s *solutionSheet) Append(lines ...interface{}) {
	f, err := os.Create(s.Name)
	check(err)

	//newText := fmt.Sprintln(lines...)
	newText := ""
	for _, line := range lines {
		if _, ok := line.(float64); ok {
			newText += fmt.Sprintf("%5.2f ", line)
			continue
		}
		newText += fmt.Sprintf("%v ", line)
	}
	newText += "\n"
	fmt.Print(newText)
	s.Body += newText

	_, err = f.WriteString(s.Body)
	check(err)

	check(f.Close())
}

func (s *solutionSheet) Title(title string) {
	s.Append(format("### %-15v ###############", title))
}
