package engine

import (
	"strconv"
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)
	if len(fields) >= 2 {
		name := fields[0]
		args := fields[1:]
		if name == "print" {
			text := strings.Join(args, " ")
			return &PrintCommand{Arg: text}
		} else if name == "add" {
			arg1, err := strconv.Atoi(args[0])
			if err != nil {
				return &PrintCommand{Arg: "SYNTAX ERROR: " + err.Error()}
			}

			arg2, err := strconv.Atoi(args[1])
			if err != nil {
				return &PrintCommand{Arg: "SYNTAX ERROR: " + err.Error()}
			}

			return &AddCommand{Arg1: arg1, Arg2: arg2}
		} else {
			return &PrintCommand{Arg: "SYNTAX ERROR: Invalid command name (" + name + ")"}
		}
	}

	return nil
}
