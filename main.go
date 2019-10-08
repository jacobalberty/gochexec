package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	c := check{}
	c.init()

	flag.BoolVar(&c.NoCSV, "nocsv", false, "turn off CSV processing")
	flag.Parse()

	fArgs := flag.Args()

	if len(fArgs) < 2 {
		fmt.Println("Not enough arguments supplied")
		fmt.Printf("Expected: %s <path> <executable>\n", os.Args[0])
		os.Exit(1)
	}

	path := fArgs[0]
	command := fArgs[1]
	params := fArgs[2:]

	defer c.ShowErrors()

	if !c.Path(path) {
		c.ErrorList = append(c.ErrorList, errors.New("could not validate specified path"))
		return
	}

	cmd := exec.Command(command, params...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		c.ErrorList = append(c.ErrorList, err)
	}
}
