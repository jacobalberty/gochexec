package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments supplied")
		fmt.Printf("Expected: %s <path> <executable>\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	command := os.Args[2]
	params := os.Args[3:]

	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("Could not access specified path.")
		os.Exit(1)
	}

	cmd := exec.Command(command, params...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}
