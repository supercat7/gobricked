package main

import (
	"gobricked/client/art"
	"gobricked/client/commands"
	"gobricked/client/shell"
)

func main() {
	baseShell := shell.NewShell(art.GOBRICKED_PROMPT, commands.SERVER_SHELL_COMMANDS)
	baseShell.RegisterCommand("exit", commands.ServerExitCommand)
	baseShell.RegisterCommand("auth", commands.ServerAuthCommand)
	// baseShell.RegisterCommand("server", commands.ServerServerCommand)
	// baseShell.RegisterCommand("listener", commands.ServerListenerCommand)
	// baseShell.RegisterCommand("agent", commands.ServerAgentCommand)
	baseShell.Start()

}
