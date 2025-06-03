package sshell

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ShellSettings struct {
	promt    string
	exitMsg  string
	commands []Command
}

type Command struct {
	input   string
	handler func() string
}

func (s ShellSettings) getCommandIndex(command string) int {

	for i, c := range s.commands {
		if c.input == command {
			return i
		}
	}

	return -1
}

func StartShell(s ShellSettings) {
	reader := bufio.NewReader(os.Stdin)
	input := ""

	for input != s.exitMsg {

		fmt.Print(s.promt)
		var err error //necessary so input is not redeclared
		input, err = reader.ReadString('\n')

		if err != nil {
			log.Fatal("error reading input", err)
		}

		// remove newline
		input = input[:len(input)-1]

		handleInput(s, input)
	}
}

func handleInput(s ShellSettings, input string) {

	if input == "" {
		return
	}

	index := s.getCommandIndex(input)
	if index == -1 {
		fmt.Println("unrecognized command")
		return
	} else {
		fmt.Println(s.commands[index].handler())
	}
}
