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

		// if i want to exit the shell
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}

		// if i want echo to my command, method hasPrefix helps to read before the argument passed
		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else {
			fmt.Println(command + ": command not found")
		}

	}
}
