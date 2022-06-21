package engine

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}
