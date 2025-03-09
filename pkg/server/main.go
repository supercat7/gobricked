package main

import (
	"fmt"
	"gobricked/pkg/server/art"
	"gobricked/pkg/server/commands"
	"gobricked/pkg/server/shell"
	"gobricked/pkg/server/stats"
)

func main() {
	fmt.Printf("%s\n", art.GOBRICKED_BANNER)

	fmt.Println("Initalizing Server statistics and data...")
	stats.UpTimeInit()

	fmt.Println("Loading Database...")

	fmt.Println("Loading Web Components and HTTP Server...")

	baseShell := shell.NewShell(art.GOBRICKED_PROMPT, commands.SERVER_SHELL_COMMANDS)
	baseShell.RegisterCommand("exit", commands.ServerExitCommand)
	baseShell.RegisterCommand("server", commands.ServerServerCommand)
	baseShell.RegisterCommand("agent", commands.ServerAgentCommand)
	baseShell.Start()
}
