/*
 * MIT License
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"./smi"
	"fmt"
	"os"
)

func main() {

	// Ref. https://gobyexample.com/command-line-arguments
	argsWithProg := os.Args

	if len(argsWithProg) > 1 {
		// os.Args[0] will be "smi-main"
		switch os.Args[1] {
		case "s":
			smi.Server()
		case "c":
			smi.Client()
		}

	} else {
		fmt.Println("Please specify the mode: s for server, c for client: smi-main s or smi-main c")
	}
}
