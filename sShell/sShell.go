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
	SingleMode     bool
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

func StartShell(s ShellSettings) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input := ""
	output := ""

	if s.DefaultHandler == nil {
		return "", errors.New("no default handler specified")
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

		output = handleInput(s, input)
		if s.SingleMode {
			return output, nil
		} else {
			fmt.Println(output)
		}
	}

	return "", nil
}

func handleInput(s ShellSettings, input string) string {

	if input == "" || input == s.ExitMsg {
		return ""
	}

	split := strings.Split(input, " ")
	cmd := split[0]
	args := split[1:]

	if cmd == "help" {
		return handleHelp(s, args)
	}

	index := s.getCommandIndex(cmd)

	if index == -1 {
		return s.DefaultHandler(split)
	} else {
		return s.Commands[index].Handler(args)
	}
}

func handleHelp(s ShellSettings, args []string) string {

	if len(args) == 0 {
		//get list of commands
		commands := ""
		for i := range s.Commands {
			commands += s.Commands[i].Input + "\n"
		}

		return "please specify what command you want help for. The available commands are:" + commands
	}

	if args[0] == "help" {
		return "help <cmd>: prints the help message for the specified command"
	}

	index := s.getCommandIndex(args[0])

	if index != -1 {
		return s.Commands[index].HelpMsg
	} else {
		return "unrecognized command"
	}
}
