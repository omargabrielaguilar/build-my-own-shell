package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var _ = fmt.Print

func main() {
	// el reader
	reader := bufio.NewReader(os.Stdin)

	// lista de builtin a usar
	commands := []string{
		"exit",
		"echo",
		"type",
	}

	// IMPORTANTE EN GO:
	// value, err := something()

	for {
		fmt.Print("$ ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}

		if strings.HasPrefix(command, "type ") {

			arg := strings.TrimSpace(command[5:])

			// en reemplazo de for, un loop manual reemplazado por slices
			if slices.Contains(commands, arg) {
				fmt.Println(arg + " is a shell builtin")
				continue
			}

			// buscar executable   usando path
			path, err := exec.LookPath(arg)

			if err == nil {
				fmt.Println(arg + " is " + path)
			} else {
				fmt.Println(arg + ": not found")
			}

			continue
		}

		fmt.Println(command + ": command not found")
	}
}
