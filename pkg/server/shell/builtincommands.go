package shell

import "fmt"

func HelpCommand(args []string, cmddesc map[string][]string) {
	if len(args) == 0 {
		keys := []string{"help", "exit", "server", "agent"}
		for i := 0; i < len(keys); i++ {
			fmt.Printf("-> %s : %s\n", keys[i], cmddesc[keys[i]][0])
		}
		fmt.Println("\nTo get more info about a command run 'help <command>'\n")
	} else {
		_, ok := cmddesc[args[0]]
		if ok {
			fmt.Printf("\n%s : %s\n", args[0], cmddesc[args[0]][1])
		} else {
			fmt.Println("Command not found: ", args[0])
		}
	}
}
