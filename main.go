package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var errorList []error

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments supplied")
		fmt.Printf("Expected: %s <path> <executable>\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	command := os.Args[2]
	params := os.Args[3:]

	defer waitCont()

	if !checkPath(path) {
		errorList = append(errorList, errors.New("Could not access specified path"))
		return
	}
	cmd := exec.Command(command, params...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		errorList = append(errorList, err)
	}
}

func waitCont() {
	if len(errorList) > 0 {
		for _, err := range errorList {
			fmt.Println(err)
		}
		fmt.Println("Press the Enter Key to terminate the console screen!")
		var input string
		fmt.Scanln(&input)
		os.Exit(1)
	}
}

func checkPath(path string) (ret bool) {
	ret = true
	_, err := os.Stat(path)
	if err != nil {
		ret = false
	}
	return
}
