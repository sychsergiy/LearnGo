package cli

import (
	"log"
	"os"
)

type Command interface {
	Execute()
}

func RetrieveArgs(argsLen int) []string {
	if len(os.Args) < argsLen+2 {
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
