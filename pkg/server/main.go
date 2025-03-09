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

	baseShell := shell.NewShell(art.GOBRICKED_PROMPT, commands.SERVER_SHELL_COMMANDS)
	baseShell.RegisterCommand("exit", commands.ServerExitCommand)
	baseShell.RegisterCommand("server", commands.ServerServerCommand)
	baseShell.RegisterCommand("agent", commands.ServerAgentCommand)
	baseShell.Start()
}
