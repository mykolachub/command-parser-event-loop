package engine

import (
	"fmt"
)

func NewAddCommand(arg1, arg2 int) *addCommand {
	return &addCommand{
		arg1: arg1,
		arg2: arg2,
	}
}

type addCommand struct {
	arg1, arg2 int
}

func (add *addCommand) Execute() {
	fmt.Println(add.arg1 + add.arg2)
	// Return it later
}
