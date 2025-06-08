# sShell - A simple golang shell library

## installation

~~~bash
go get github.com/theskibb/sShell
~~~

you can now import the library with 

~~~go
go get github.com/theskibb/sShell
~~~

## how to use

to create a shell using the library you need to create a **ShellSettings** struct

**ShellSettings** needs three fields:

- **Prompt**: the prompt indicator for the shell
- **ExitMsg**: typing this message will stop the shell
- **Commands**: an array of **Command** structs
- **DefaultHandler**: handles inputs which are not in the **commands** array
    - needs to be a func(args []string) string

**Command** structs need three fields:

- **Input**: the command name you write into the shell
- **handler**: the function which will execute when you run the command
    - needs to be a func(args []string) string
- **HelpMsg**: the string which will be displayed for this command when you run **help**

## examples

for examples of how to use the library, see the examples folder.


