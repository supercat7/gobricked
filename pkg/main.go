package main

import (
	"fmt"
	"gobricked/pkg/art"
	"gobricked/pkg/commands"
	"gobricked/pkg/comms"
	"gobricked/pkg/shell"
	"gobricked/pkg/stats"
)

func main() {
	fmt.Printf("%s\n", art.GOBRICKED_BANNER)

	fmt.Println("Initalizing Server statistics and data...")
	stats.UpTimeInit()

	fmt.Println("Loading server configurations...")

	fmt.Println("Launching TCP Server on port:", "9090")
	go comms.SERVERINSTANCE.Start(comms.SERVERCHANNEL)

	fmt.Println("Loading Database...")

	fmt.Println("Loading Web Components and HTTP Server...")

	baseShell := shell.NewShell(art.GOBRICKED_PROMPT, commands.SERVER_SHELL_COMMANDS)
	baseShell.RegisterCommand("exit", commands.ServerExitCommand)
	baseShell.RegisterCommand("server", commands.ServerServerCommand)
	baseShell.RegisterCommand("agent", commands.ServerAgentCommand)
	baseShell.Start()
}
