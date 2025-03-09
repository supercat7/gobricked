package commands

import (
	"fmt"
)

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":   {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":   {"Exits gobricked", "Exits gobricked"},
	"server": {"Server actions", "Options: server <start|info|stop> <port>\n-> Leave <port> empty if you want to shutdown all server instances\n-> for option <info> you don't need to enter a port"},
	"agent":  {"Agent actions", "Options: agent <pick|list> <AGENT ID>\n-> <pick> Spawns a subshell to control agent\n->Leave <AGENT_ID> empty for <list> option"},
}

func ExitCommand(args []string) {
	fmt.Printf("Shutting down Server instances...\n")
}
