package engine

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(c Command)
}

type EventLoop struct {
	q *CommandQueue

	stopSignal chan struct{}
	stop bool
}

// Start executing commands
func (l *EventLoop) Start() {
	l.q = &CommandQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stop || !l.q.empty(){
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

// Add Command to the Command Queue
func (l *EventLoop) Post(c Command) {
	l.q.push(c)
}

// Await until all commands executed
func (l *EventLoop) AwaitFinish() {
}

type CommandQueue struct {
}

// Push command to the Queue
func (cq *CommandQueue) push(c Command) {

}

// Pull command from the Queue
func (cq *CommandQueue) pull(c Command) {

}

// Is Command Queue empty
func (cq *CommandQueue) isEmpty(c Command) {

}
