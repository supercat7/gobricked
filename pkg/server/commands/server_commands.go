package commands

import (
	"fmt"
)

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help": {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit": {"Exits gobricked", "Exits gobricked"},
}

func ListCommands(commands map[string][]string) {
	for key, value := range commands {
		fmt.Printf("-> %s : %s\n", key, value[0])
	}
	fmt.Println("To get more info about a command run 'help <command>'")
}

func ServerHelp(input []string) {
	if len(input) == 1 {
		ListCommands(SERVER_SHELL_COMMANDS)
	} else if len(input) > 1 {
		_, ok := SERVER_SHELL_COMMANDS[input[1]]
		if ok {
			fmt.Printf("-> %s : %s\n", input[1], SERVER_SHELL_COMMANDS[input[1]][1])
		} else {
			fmt.Printf("No such command %s...\n", input[1])
		}
	}
}
