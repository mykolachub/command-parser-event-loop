package engine

import (
	"strconv"
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)

	if len(fields) < 2 {
		return NewPrintCommand("SYNTAX ERROR: Not Enough Parameters")
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "print":
		message := strings.Join(args, " ")
		return NewPrintCommand(message)
	case "add":
		arg1, err := strconv.Atoi(args[0])
		if err != nil {
			return NewPrintCommand("SYNTAX ERROR: " + err.Error())
		}

		arg2, err := strconv.Atoi(args[1])
		if err != nil {
			return NewPrintCommand("SYNTAX ERROR: " + err.Error())
		}

		return NewAddCommand(arg1, arg2)
	default:
		return NewPrintCommand("ERROR: Unknown Command (" + in + ")")
	}
}
