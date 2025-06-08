package main

import (
	"fmt"
	"github.com/theskibb/sShell/sShell"
)

func main() {
	fmt.Println()

	s := sshell.ShellSettings{
		Promt:   ">",
		ExitMsg: "exit",
		Commands: []sshell.Command{
			sshell.Command{
				Input:   "rev",
				Handler: reverse,
				HelpMsg: "reverses the order of the words input",
			},
		},
		DefaultHandler: func(args []string) string { return "unrecognized command" },
		SingleMode:     true,
	}

	result, err := sshell.StartShell(s)

	if err != nil {
		fmt.Println("something went wrong", err)
	}

	fmt.Println("outputting result of shell operation from outside the shell:", result)
}

func reverse(args []string) string {
	output := ""
	for i := len(args) - 1; i >= 0; i-- {
		output += args[i] + " "
	}

	return output
}
