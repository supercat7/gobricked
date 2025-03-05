package shell

import "fmt"

var SHELL_COMMANDS = map[string]string{
	"help": "Prints commands",
	"exit": "Exits gobricked",
}

func ListCommands() {
	for key, value := range SHELL_COMMANDS {
		fmt.Printf("-> %s : %s\n", key, value)
	}
}
