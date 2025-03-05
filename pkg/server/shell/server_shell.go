package shell

import (
	"bufio"
	"fmt"
	"gobricked/pkg/server/commands"
	"os"
	"strings"
)

func ParseArgs() []string {
	fmt.Printf("gobricked> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	inputArr := strings.Fields(input)

	return inputArr
}

func ParseCommands(input []string, shellCommands map[string][]string) {
	if input[0] == "help" {
		commands.Help(input, shellCommands)
	} else if input[0] == "exit" {
		commands.ServerExit()
	} else if input[0] == "" {
		//
	} else {
		fmt.Println("No such command exists use command 'help' for more info")
	}
}

func Shell(shellCommands map[string][]string) {
	for {
		input := ParseArgs()
		ParseCommands(input, shellCommands)
	}
}

// make this shell an object so I can pass in args as needed and spawn a new shell?
