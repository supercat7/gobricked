package commands

import "fmt"

var AGENT_SHELL_COMMANDS = map[string][]string{
	"help": {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit": {"Exits Agent subshell", "Exits Agent subshell"},
}

func AgentHelpCommand(args []string) {
	fmt.Println("Agent Help...")
}

func AgentExitCommand(args []string) {

}
