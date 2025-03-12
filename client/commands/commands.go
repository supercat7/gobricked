package commands

import (
	"gobricked/client/comms"
	"os"
	"strings"
)

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":     {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":     {"Exits gobricked shell", "Exits gobricked shell"},
	"auth":     {"Authenticates and connects to server", "Usage: auth 192.168.0.1:9090 username:password"},
	"server":   {"Server actions", "Options: server <info>\n-> Displays information about server instance"},
	"listener": {"Listener for agents", "Options: listener <port>"},
	"agent":    {"Agent actions", "Options: agent <pick|info|list> <AGENT ID>\n-> <pick> Spawns a subshell to control agent\n->Leave <AGENT_ID> empty for <list> option"},
}

func ServerExitCommand(args []string) {
	os.Exit(0)
}
func ServerAuthCommand(args []string) {
	dst := strings.Split(args[0], ":")
	userPass := strings.Split(args[1], ":")

	s := comms.SockStream{
		IP:   dst[0],
		Port: dst[1],
	}
	s.AuthServer(userPass[0], userPass[1])
}
