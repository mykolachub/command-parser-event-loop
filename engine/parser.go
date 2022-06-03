package engine

import (
	"strconv"
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)

	if len(fields) < 2 {
		return &PrintCommand{Arg: "SYNTAX ERROR: Not Enough Parameters"}
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "print":
		message := strings.Join(args, " ")
		return &PrintCommand{Arg: message}
	case "add":
		arg1, err := strconv.Atoi(args[0])
		if err != nil {
			return &PrintCommand{Arg: "SYNTAX ERROR: " + err.Error()}
		}

		arg2, err := strconv.Atoi(args[1])
		if err != nil {
			return &PrintCommand{Arg: "SYNTAX ERROR: " + err.Error()}
		}

		return &AddCommand{Arg1: arg1, Arg2: arg2}
	default:
		return &PrintCommand{Arg: "ERROR: Unkown Command (" + in + ")"}
	}
}
