package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Shell struct {
	Prompt   string
	Commands map[string]func([]string)
	CmdDesc  map[string][]string
}

func (s *Shell) RegisterCommand(cmd string, handler func([]string)) {
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Println("\nReceived signal:", sig)
		fmt.Println("Exiting shell...")
		os.Exit(0)
	}()

	for {
		fmt.Print(s.Prompt)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		input := scanner.Text()
		inputArr := strings.Fields(input)

		if len(inputArr) > 0 {
			cmd, args := inputArr[0], inputArr[1:]
			if cmd == "help" {
				HelpCommand(args, s.CmdDesc)
				continue
			}
			if handler, exists := s.Commands[cmd]; exists {
				handler(args)
				if cmd == "exit" {
					break
				}
			} else {
				fmt.Println("Command not found: ", cmd)
			}
		}
	}
}
