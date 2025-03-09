package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	Prompt   string
	Commands map[string]func([]string)
	CmdDesc  map[string][]string
}

func (s *Shell) AddCommand(cmd string, handler func([]string)) {
	s.Commands[cmd] = handler
}

func NewShell(prompt string, cmddesc map[string][]string) *Shell {
	return &Shell{
		Prompt:   prompt,
		Commands: make(map[string]func([]string)),
		CmdDesc:  cmddesc,
	}
}

func (s *Shell) Start() {
	for {
		fmt.Printf(s.Prompt)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := scanner.Text()
		inputArr := strings.Fields(input)

		if len(inputArr) > 0 {
			cmd, args := inputArr[0], inputArr[1:]
			if handler, exists := s.Commands[cmd]; exists {
				handler(args)
			} else {
				fmt.Println("Command not found: ", cmd)
			}
		}
	}
}
