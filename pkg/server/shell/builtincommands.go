package shell

import (
	"fmt"
	"gobricked/pkg/server/util"
)

func HelpCommand(args []string, cmddesc map[string][]string) {
	if len(args) == 0 {
		for key, value := range cmddesc {
			fmt.Printf("%s : %s\n", key, value[0])
		}
		fmt.Println("\nTo get more info about a command run 'help <command>'\n")
	} else {
		_, ok := cmddesc[args[0]]
		if ok {
			fmt.Printf("\n%s : %s\n", args[0], cmddesc[args[0]][1])
		} else {
			util.CommandNotFoundErr(args[0])
		}
	}
}
