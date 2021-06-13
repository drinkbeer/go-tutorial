package cli

import (
	"fmt"

	commands "github.com/drinkbeer/go-tutorial/cli/commands"
)

// Command represents a specific command in a command hierarchy.
type Command interface {
	// Name is the label used to invoke the command.
	Name() string
	// Usage prints information about the command syntax and any additional explanation.
	Usage()
	// Description is an one-line description of the purpose of the command.
	Description() string
	// Run receives the CLI arguments following the command name and executes the command.
	Run(args ...string) error
}

type Commands []Command

type CLI struct {
	commands Commands
}

func NewCLI() *CLI {
	c := &CLI{}
	c.commands = append(c.commands, commands.NewPalindromeCheck())
	return c
}

func (cli *CLI) RunCommand(args ...string) error {
	fmt.Println("Start running command", args[0])
	for _, c := range cli.commands {
		if c.Name() == args[0] {
			return c.Run(args[1:]...)
		}
	}
	return fmt.Errorf("Unknown command '%s', supported commands: %s", args[0], cli.commands)
}
