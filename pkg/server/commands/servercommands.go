package commands

import (
	"fmt"
	"gobricked/pkg/server/shell"
	"os"
)

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":   {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":   {"Exits gobricked", "Exits gobricked"},
	"server": {"Server actions", "Options: server <start|info|stop> <port>\n-> leave <port> empty if you want to shutdown all server instances\n-> for option <info> you don't need to enter a port"},
	"agent":  {"Agent actions", "Options: agent <pick|info|list> <AGENT ID>\n-> leave <AGENT_ID> empty for <list> option"},
}

func (s *shell.Shell) HelpCommand(args []string) {
	if len(args) == 0 {
		keys := []string{"help", "exit", "server", "agent"}
		for i := 0; i < len(keys); i++ {
			fmt.Printf("-> %s : %s\n", keys[i], SERVER_SHELL_COMMANDS[keys[i]][0])
		}
		fmt.Println("\nTo get more info about a command run 'help <command>'\n")
	} else {
		_, ok := SERVER_SHELL_COMMANDS[args[0]]
		if ok {
			fmt.Printf("\n%s : %s\n", args[0], SERVER_SHELL_COMMANDS[args[0]][1])
		} else {
			fmt.Println("Command not found: ", args[0])
		}
	}
}

func ExitCommand(args []string) {
	os.Exit(0)
}

func ServerServerCommand(args []string) {}
