package commands

import (
	"fmt"
	"gobricked/pkg/server/art"
	"gobricked/pkg/server/comms"
	"gobricked/pkg/server/shell"
	"gobricked/pkg/server/stats"
	"strconv"
)

var SERVER_SHELL_COMMANDS = map[string][]string{
	"help":   {"Prints commands", "Options: help <command>	Lists more information about a command"},
	"exit":   {"Exits gobricked", "Exits gobricked"},
	"server": {"Server actions", "Options: server <info>\n-> Displays information about server instance"},
	"agent":  {"Agent actions", "Options: agent <pick|info|list> <AGENT ID>\n-> <pick> Spawns a subshell to control agent\n->Leave <AGENT_ID> empty for <list> option"},
}

func ServerExitCommand(args []string) {
	comms.SERVERINSTANCE.Stop(comms.SERVERCHANNEL)
}

func ServerServerCommand(args []string) {
	if len(args) > 0 {
		if args[0] == "info" {
			stats.DisplayServerStats()
		} else if args[0] == "list" {

		} else {
			fmt.Println("Err: Command not found: ", args[0])
		}
	} else {
		fmt.Println("No arguments found, use help <command>")
	}
}

func ServerAgentCommand(args []string) {
	if len(args) > 0 {
		if args[0] == "pick" {
			if len(args) == 2 {
				agentID, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Printf("Not a valid Agent ID, must be an integer: %v\n", err)
					return
				}
				prompt := fmt.Sprintf(art.GOBRICKED_AGENT_PROMPT, agentID)
				agentShell := shell.NewShell(prompt, AGENT_SHELL_COMMANDS)
				agentShell.RegisterCommand("exit", AgentExitCommand)
				agentShell.Start()
			} else {
				fmt.Println("No arguments found, use help <command>")
			}
		} else if args[0] == "list" {
			// list agents
		} else {
			fmt.Println("Err: Command not found: ", args[0])
		}
	} else {
		fmt.Println("No arguments found, use help <command>")
	}
}
