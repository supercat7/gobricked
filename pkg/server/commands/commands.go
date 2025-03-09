package commands

import (
	"os"
	"fmt"
)
func HelpCommand(args []string) {
	if len(args) == 0 {
		for key, value := range SERVER_SHELL_COMMANDS {
			fmt.Printf("-> %s : %s\n", key, value[0])
		}
		fmt.Println("To get more info about a command run 'help <command>'")
	}
}


var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":   {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":   {"Exits gobricked", "Exits gobricked"},
	"server": {"Server actions", "Options: server <start|info|stop> <port>\n-> leave <port> empty if you want to shutdown all server instances\n-> for option <info> you don't need to enter a port"},
}

func ServerExit() {
	os.Exit(0)
}
