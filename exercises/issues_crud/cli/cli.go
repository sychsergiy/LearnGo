package cli

import (
	"log"
	"os"
)

type cli struct {
	handlers map[string]Handler
}

func New() *cli {
	instance := cli{}
	instance.handlers = make(map[string]Handler)
	return &instance
}

func (c *cli) RegisterHandler(key string, handler Handler) {
	c.handlers[key] = handler
}

func (c *cli) handle(key string) {
	if actionHandler, ok := c.handlers[key]; ok {
		actionHandler.Execute()
	} else {
		log.Fatal("No registered handler on action: " + key, )
	}
}

func (c *cli) Run() {
	action := os.Args[1]
	c.handle(action)
}
