package engine

import "sync"

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
	mu sync.Mutex
	a  []Command

	notEmpty chan struct{}
	wait     bool
}

// Push command to the Queue
func (cq *CommandQueue) push(c Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.a = append(cq.a, c)
	if cq.wait {
		cq.notEmpty <- struct{}{}
	}
}

// Pull command from the Queue
func (cq *CommandQueue) pull(c Command) {

}

// Is Command Queue empty
func (cq *CommandQueue) isEmpty(c Command) {

}
