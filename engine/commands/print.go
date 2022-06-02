package engine

import (
	"fmt"
)

func NewPrintCommand(arg string) *printCommand {
	return &printCommand{
		arg: arg,
	}
}

type printCommand struct {
	arg string
}

func (p *printCommand) Execute() {
	fmt.Println(p.arg)
}
