package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt"
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, _ := reader.ReadString('\n')
		builtinToCheck := command[5:]

		// 1. if i want to exit the shell
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}

		// 2.  if i want echo to my command, method hasPrefix helps to read before the argument passed
		if strings.HasPrefix(command, "echo ") {
			fmt.Println(builtinToCheck)
		} else {
			fmt.Println(command + ": command not found")
		}

		// slice mode
		commands := []string{
			"exit",
			"echo",
			"type",
		}

		// 3. Work with typeee
		if strings.HasPrefix(command, "type ") {
			for _, builtinToCheck := range commands {
				fmt.Println(builtinToCheck + "is a shell builtin")
			}
		}
	}
}
