/*
 * This main function essentially launches in two different modes based on
 * command line arguments passed, If you pass argument 'c', the process runs as
 * the client of the server and if you pass the argument as 's' it runs in the
 * server mode. Naturally, you will need to start the server process first for
 * any client connection to get established.
 *
 * MIT License
 * Author: Umesh Patil, Neosemantix, Inc.
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
