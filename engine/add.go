package engine

import "strconv"

type AddCommand struct {
	Arg1, Arg2 int
}

func (add *AddCommand) Execute(h Handler) {
	res := add.Arg1 + add.Arg2
	h.Post(&PrintCommand{Arg: strconv.Itoa(res)})
}
