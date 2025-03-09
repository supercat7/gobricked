package commands

import (
	"fmt"
	"gobricked/pkg/server/art"
	"gobricked/pkg/server/shell"
	"gobricked/pkg/server/util"
	"strconv"
)

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":   {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":   {"Exits gobricked", "Exits gobricked"},
	"server": {"Server actions", "Options: server <start|info|stop> <port>\n-> Leave <port> empty if you want to shutdown all server instances\n-> for option <info> you don't need to enter a port"},
	"agent":  {"Agent actions", "Options: agent <pick|list> <AGENT ID>\n-> <pick> Spawns a subshell to control agent\n->Leave <AGENT_ID> empty for <list> option"},
}

func ServerExitCommand(args []string) {
	fmt.Printf("Shutting down Server instances...\n")
}

func ServerServerCommand(args []string) {

}

func ServerAgentCommand(args []string) {
	if len(args) > 0 {
		if args[0] == "pick" {
			agentID, _ := strconv.Atoi(args[1])
			prompt := fmt.Sprintf(art.GOBRICKED_AGENT_PROMPT, agentID)
			agentShell := shell.NewShell(prompt, AGENT_SHELL_COMMANDS)
			agentShell.RegisterCommand("exit", AgentExitCommand)
			agentShell.Start()
		}
	} else {
		util.CommandNoArgumentErr()
	}
}
