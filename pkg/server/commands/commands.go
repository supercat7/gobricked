package commands

import (
	"fmt"
)

var SHELL_COMMANDS = map[string]string{
	"help": "Prints commands",
	"exit": "Exits gobricked",
}

var SUBSHELL_COMMANDS = map[string]string{
	"help": "Options: help <command>	Lists more information about a command",
}

func ListCommands() {
	for key, value := range SHELL_COMMANDS {
		fmt.Printf("-> %s : %s\n", key, value)
	}
	fmt.Println("To get more info about a command run 'help <command>'")
}

func Help(input []string) {
	if len(input) == 1 {
		ListCommands()
	} else if len(input) > 1 {
		_, ok := SUBSHELL_COMMANDS[input[1]]
		if ok {
			fmt.Printf("%s : %s\n", input[0], SUBSHELL_COMMANDS[input[1]])
		} else {
			fmt.Printf("No such command %s...\n", input[1])
		}
	}
}
