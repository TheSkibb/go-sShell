package main

import (
	"fmt"
	"github.com/theskibb/sShell/sShell"
	"strconv"
)

func main() {
	fmt.Println()

	s := sshell.ShellSettings{
		Promt:   ">",
		ExitMsg: "exit",
		Commands: []sshell.Command{
			{
				Input:   "add",
				Handler: add,
				HelpMsg: "add <arg1> <arg2> <arg3> ...",
			},
			{
				Input:   "sub",
				Handler: subtract,
				HelpMsg: "sub <arg1> <arg2> <arg3> ...",
			},
		},
	}

	sshell.StartShell(s)
}

func add(args []string) string {

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

func subtract(args []string) string {

	total := 0

	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return arg + "is not a number"
		}
		if i == 0 {
			total = num
		} else {
			total -= num
		}

	}

	totalStr := strconv.Itoa(total)

	return totalStr
}

/*
func calc(args []string) string {

}
*/
