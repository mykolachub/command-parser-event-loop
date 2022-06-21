package engine

import (
	"sync"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(c Command)
}

type EventLoop struct {
	q *commandQueue

	stopSignal chan struct{}
	stop       bool
}

// Start executing commands
func (l *EventLoop) Start() {
	l.q = &commandQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	l.stop = false
	go func() {
		for !l.stop || !l.q.empty() {
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
	l.Post(&stopCommand{})
	<-l.stopSignal
}

type commandQueue struct {
	mu sync.Mutex
	a  []Command

	notEmpty chan struct{}
	wait     bool
}

// Push command to the Queue
func (cq *commandQueue) push(c Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.a = append(cq.a, c)

	if cq.wait {
		cq.wait = false
		cq.notEmpty <- struct{}{}
	}
}

// Pull command from the Queue
func (cq *commandQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.empty() {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}

	res := cq.a[0]
	cq.a[0] = nil
	cq.a = cq.a[1:]
	return res
}

// Is Command Queue empty
func (cq *commandQueue) empty() bool {
	return len(cq.a) == 0
}
