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

	_, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println("Could not access specified path.")
		os.Exit(1)
	}
	fmt.Println(os.Args[3:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}
