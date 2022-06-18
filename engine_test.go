package main

import (
	"github.com/nikolaichub/command-parser-event-loop/engine"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type addCommand struct {
	arg1, arg2 int
}

// Must return add command
func TestParserAdd(t *testing.T) {
	commandString := "add 1 2"
	command := engine.Parse(commandString)
	exampleAdd := engine.NewAddCommand(1, 2)

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(exampleAdd), reflect.TypeOf(command))
	}
}

// Must return print command
func TestParserPrint(t *testing.T) {
	commandString := "print Hello world"
	command := engine.Parse(commandString)
	examplePrint := engine.NewPrintCommand("Hello world")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}

// Must return print command (containing error message)
func TestParserDefault(t *testing.T) {
	commandString := ""
	command := engine.Parse(commandString)
	examplePrint := engine.NewPrintCommand("Hello world")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}
