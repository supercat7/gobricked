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

func ParseCommands(input []string) {
	if input[0] == "help" {
		commands.Help(input)
	} else if input[0] == "exit" {
		os.Exit(0)
	} else if input[0] == "" {
		//
	} else {
		fmt.Println("No such command exists use command 'help' for more info")
	}
}

func Shell() {
	for {
		input := ParseArgs()
		ParseCommands(input)
	}
}
