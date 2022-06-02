package engine

import (
	"fmt"
	engine "github.com/command-parser-event-loop/engine/commands"
	"strconv"
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)
	if len(fields) > 2 {
		name := fields[0]
		args := fields[1:]
		if name == "print" {
			text := strings.Join(args, " ")
			return engine.NewPrintCommand(text)
		} else if name == "add" {
			arg1, err := strconv.Atoi(args[0])
			if err != nil {
				return engine.NewPrintCommand(fmt.Sprintf("SYNTAX ERROR: %s", err.Error()))
			}

			arg2, err := strconv.Atoi(args[1])
			if err != nil {
				return engine.NewPrintCommand(fmt.Sprintf("SYNTAX ERROR: %s", err.Error()))
			}

			return engine.NewAddCommand(arg1, arg2)
		} else {
			return engine.NewPrintCommand(fmt.Sprintf("SYNTAX ERROR: Invalid command name (%s)", name))
		}
	}

	return nil
}
