package cmd

import "os"

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":     {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":     {"Exits gobricked shell", "Exits gobricked shell"},
	"server":   {"Server actions", "Options: server <info>\n-> Displays information about server instance"},
	"listener": {"Listener for agents", "Options: listener <port>"},
	"agent":    {"Agent actions", "Options: agent <pick|info|list> <AGENT ID>\n-> <pick> Spawns a subshell to control agent\n->Leave <AGENT_ID> empty for <list> option"},
}

var AGENT_SHELL_COMMANDS = map[string][]string{
	"help": {"Prints commands", "Options: help <command>	Lists more information about a command"},
}

func ServerExitCommand(args []string) {
	os.Exit(0)
}
