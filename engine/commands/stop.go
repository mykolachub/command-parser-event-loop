package engine

func NewStopCommand(arg string) *stopCommand {
	return &stopCommand{}
}

type stopCommand struct {}

func (s stopCommand) Execute(h Handler){
	h.(*Loop).stop = true
}
