package engine

import (
	"strconv"
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

func (add *addCommand) Execute(h Handler) {
	res := add.arg1 + add.arg2
	h.Post(NewPrintCommand(strconv.Itoa(res)))
}
