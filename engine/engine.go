package engine

type Command interface {
}

type Handler interface {
}

type EventLoop struct {
}

// Start executing commands
func (l *EventLoop) Start() {

}

// Add Command to the Command Queue
func (l *EventLoop) Post(c Command) {

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
