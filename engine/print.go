package engine

import "fmt"

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(h Handler) {
	fmt.Println(p.Arg)
}
