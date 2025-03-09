package util

import "fmt"

func CommandNotFoundErr(cmd string) {
	fmt.Println("Err: Command not found: ", cmd)
}

func CommandNoArgumentErr() {
	fmt.Println("Err: No argument provided, use help <command>")
}
