package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type check struct {
	NoCSV     bool
	cPaths    map[string]bool
	ErrorList []error
}

func (c *check) init() {
	c.NoCSV = false
	c.cPaths = make(map[string]bool)
}

func (c *check) Path(path string) bool {
	// Prevent the same file from being parsed twice
	if val, ok := c.cPaths[path]; ok {
		return val
	}

	c.cPaths[path] = false

	stat, err := os.Stat(path)

	if err != nil {
		return false
	}

	c.cPaths[path] = true
	// If path is a csv file then read it for more resources to check
	if !c.NoCSV && !stat.IsDir() && filepath.Ext(path) == ".csv" {
		fileR, _ := os.Open(path)
		csvR := csv.NewReader(fileR)
		for {
			record, err := csvR.Read()
			if err == io.EOF {
				break
			}

			switch record[0] {
			case "dns":
				if !c.DNS(record[1:]) {
					c.ErrorList = append(c.ErrorList, fmt.Errorf("Could not resolve host '%s'", record[1]))
				}
			case "file":
				if !c.Path(record[1]) {
					c.ErrorList = append(c.ErrorList, fmt.Errorf("Could not validate path: '%s' from '%s'", record[1], path))
				}
			case "sock":
				if !c.Sock(record[1:]) {
					c.ErrorList = append(c.ErrorList, fmt.Errorf("Could not establish connection to %s socket '%s'", record[1], record[2]))
				}
			}
		}
		if len(c.ErrorList) > 0 {
			return false
		}
	}
	return true
}

func (c *check) DNS(options []string) bool {
	r := net.Resolver{}
	ctx := context.Background()
	_, err := r.LookupIPAddr(ctx, options[0])
	if err != nil {
		return false
	}
	return true
}

func (c *check) Sock(options []string) bool {
	var timeout int = 1

	if len(options) > 2 {
		i, err := strconv.Atoi(options[2])
		if err != nil {
			timeout = i
		}
	}

	conn, err := net.DialTimeout(options[0], options[1], time.Duration(timeout)*time.Second)

	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func (c *check) ShowErrors() {
	if len(c.ErrorList) > 0 {
		for _, err := range c.ErrorList {
			fmt.Println(err)
		}
		fmt.Println("Press the Enter Key to terminate the console screen!")
		var input string
		fmt.Scanln(&input)
		os.Exit(1)
	}
}
