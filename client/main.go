package main

import (
	"fmt"
	"gobricked/client/art"
	"gobricked/client/cmd"
	"gobricked/client/comms"
	"gobricked/client/shell"
)

func main() {
	var pass string
	var remoteAdd string
	var ok bool = false

	fmt.Printf("Enter server address (Format: '127.0.0.1:9090'): ")
	fmt.Scanln(&remoteAdd)
	sock := comms.NewSockStream(remoteAdd)

	for !ok {
		fmt.Printf("Enter server password: ")
		fmt.Scanf("%s", &pass)
		ok = sock.AuthServer(pass)
	}

	baseShell := shell.NewShell(art.GOBRICKED_PROMPT, cmd.SERVER_SHELL_COMMANDS)
	baseShell.RegisterCommand("exit", cmd.ServerExitCommand)
	baseShell.Start()
}
