package cli

import (
	"log"
	"os"
)

type cli struct {
	commands map[string]Command
}

func New() *cli {
	instance := cli{}
	instance.commands = make(map[string]Command)
	return &instance
}

func (c *cli) RegisterCommand(key string, handler Command) {
	c.commands[key] = handler
}

func (c *cli) handle(key string) {
	if actionHandler, ok := c.commands[key]; ok {
		actionHandler.Execute()
	} else {
		log.Fatal("No registered handler on action: " + key, )
	}
}

func (c *cli) Run() {
	action := os.Args[1]
	c.handle(action)
}
