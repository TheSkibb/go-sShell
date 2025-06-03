package sshell

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type ShellSettings struct {
	Promt          string
	ExitMsg        string
	Commands       []Command
	DefaultHandler func(args []string) string
}

type Command struct {
	Input   string
	Handler func(args []string) string
	HelpMsg string
}

func (s ShellSettings) getCommandIndex(command string) int {

	for i, c := range s.Commands {
		if c.Input == command {
			return i
		}
	}

	return -1
}

func StartShell(s ShellSettings) error {
	reader := bufio.NewReader(os.Stdin)
	input := ""

	if s.DefaultHandler == nil {
		return errors.New("no default handler specified")
	}

	for input != s.ExitMsg {

		fmt.Print(s.Promt)
		var err error //necessary so input is not redeclared
		input, err = reader.ReadString('\n')

		if err != nil {
			log.Fatal("error reading input", err)
		}

		// remove newline
		input = input[:len(input)-1]

		handleInput(s, input)
	}

	return nil
}

func handleInput(s ShellSettings, input string) {

	if input == "" || input == s.ExitMsg {
		return
	}

	split := strings.Split(input, " ")
	cmd := split[0]
	args := split[1:]

	if cmd == "help" {
		handleHelp(s, args)
		return
	}

	index := s.getCommandIndex(cmd)

	if index == -1 {
		fmt.Println(s.DefaultHandler(split))
		return
	} else {
		fmt.Println(s.Commands[index].Handler(args))
	}
}

func handleHelp(s ShellSettings, args []string) {

	if len(args) == 0 {
		fmt.Println("please specify what command you want help for")
		return
	}

	if args[0] == "help" {
		fmt.Println("help <cmd>: prints the help message for the specified command")
		return
	}

	index := s.getCommandIndex(args[0])

	if index != -1 {
		fmt.Println(s.Commands[index].HelpMsg)
		return
	} else {
		fmt.Println("unrecognized command")
	}
}
