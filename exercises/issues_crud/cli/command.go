package cli

import (
	"log"
	"os"
)

type Command interface {
	Execute()
}

type BaseCommand struct {
}

func (command *BaseCommand) RetrieveArgs(argsLen int) []string {
	if len(os.Args) < argsLen {
		log.Fatal("Not enough arguments")
	}
	args := make([]string, 0, argsLen)

	for index, value := range os.Args {
		if index > 1 {
			args = append(args, value)
		}
	}
	return args
}
