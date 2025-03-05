package commands

import "fmt"

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
