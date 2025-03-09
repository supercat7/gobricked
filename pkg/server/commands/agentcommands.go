package commands

import "fmt"

var AGENT_SHELL_COMMANDS = map[string][]string{
	"help": {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit": {"Exits Agent subshell", "Exits Agent subshell"},
}

func AgentExitCommand(args []string) {
	fmt.Println("Stopping socket io with agent...")
}
