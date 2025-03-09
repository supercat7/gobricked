package util

import "fmt"

func CommandNotFoundErr(cmd string) {
	fmt.Println("Command not found: ", cmd)
}

func CommandNoArgumentErr() {
	fmt.Println("No argument provided, use help <command>")
}
