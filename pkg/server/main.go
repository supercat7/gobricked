package main

import (
	"fmt"
	"gobricked/pkg/server/art"
	"gobricked/pkg/server/shell"
	"strings"
)

func main() {
	fmt.Printf("%s\n", art.GOBRICKED_BANNER)
	for {
		var arg string = ""
		fmt.Printf("gobricked> ")
		fmt.Scanln(&arg)
		argSlice := strings.Split(arg, " ")

		if argSlice[0] == "help" {
			shell.ListCommands()
		} else if argSlice[0] == "exit" {
			break
		} else if argSlice[0] == "" {
			continue
		} else {
			fmt.Println("No such command exists use command 'help' for more info")
		}
	}
}
