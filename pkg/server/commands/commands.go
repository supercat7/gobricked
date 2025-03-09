package commands

import (
	"os"
	"fmt"
)
func Help(input []string, shellCommands map[string][]string) {
	if len(input) == 1 {
		ListCommands(shellCommands)
	} else if len(input) > 1 {
		_, ok := shellCommands[input[1]]
		if ok {
			fmt.Printf("-> %s : %s\n", input[1], SERVER_SHELL_COMMANDS[input[1]][1])
		} else {
			fmt.Printf("No such command %s...\n", input[1])
		}
	}
}

func ListCommands(commands map[string][]string) {
	for key, value := range commands {
		fmt.Printf("-> %s : %s\n", key, value[0])
	}
	fmt.Println("To get more info about a command run 'help <command>'")
}


var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":   {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":   {"Exits gobricked", "Exits gobricked"},
	"server": {"Server actions", "Options: server <start|info|stop> <port>\n-> leave <port> empty if you want to shutdown all server instances\n-> for option <info> you don't need to enter a port"},
}

func ServerExit() {
	os.Exit(0)
}
