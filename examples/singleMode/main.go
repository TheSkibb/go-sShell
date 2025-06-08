/*
sometimes, you only need to send one command, and get the result back
in this case, you can use singleMode
*/
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
		DefaultHandler: func(args []string) string { return "unrecognized command" }, //in single mode, unrecognized command, or help will reprompt you
		SingleMode:     true,                                                         //**single mode** is on
	}

	//the output from the command you run will end up in the result variable
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
