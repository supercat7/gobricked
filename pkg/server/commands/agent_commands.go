package commands

var AGENT_SHELL_COMMANDS = map[string][]string{
	"help": {"Prints commands", "Options: help <command> Lists more information about a command"},
	"exit": {"Exits agent shell", "Exits agent shell"},
}
