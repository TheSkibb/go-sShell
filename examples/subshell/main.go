/*
one of the things you might want to do is to start a subshell
in this example, when running the add command, a subshell will launch
*/

package main

import (
	"github.com/theskibb/sShell/sShell"
	"strconv"
)

func main() {
	s := sshell.ShellSettings{
		Promt:   ">",
		ExitMsg: "exit",
		Commands: []sshell.Command{
			{
				Input:   "add",
				Handler: add,
				HelpMsg: "add <arg1> <arg2> <arg3> ...",
			},
		},
		DefaultHandler: func(args []string) string { return "unrecognized command" },
	}

	sshell.StartShell(s)
}

func add(args []string) string {
	s := sshell.ShellSettings{
		Promt:          "(add) >",
		ExitMsg:        "exit",
		Commands:       []sshell.Command{},
		DefaultHandler: subAdd,
	}

	sshell.StartShell(s)
	return ""
}

func subAdd(args []string) string {

	total := 0

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return arg + "is not a number"
		}

		total += num
	}

	totalStr := strconv.Itoa(total)

	return totalStr
}
