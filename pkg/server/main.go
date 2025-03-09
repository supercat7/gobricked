package main

import (
	"fmt"
	"gobricked/pkg/server/art"
	"gobricked/pkg/server/commands"
	"gobricked/pkg/server/shell"
)

func main() {
	fmt.Printf("%s\n", art.GOBRICKED_BANNER)
	fmt.Println("Initializing TCP Server...")
	fmt.Println("Loading SQL Database...")
	fmt.Println("Loading Web Components and HTTP Server...")

	baseShell := shell.NewShell("gobricked> ", commands.SERVER_SHELL_COMMANDS)
	baseShell.AddCommand("help", commands.HelpCommand)
	baseShell.AddCommand("exit", commands.ExitCommand)
	baseShell.Start()
}
