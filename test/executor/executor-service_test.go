// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"../../executor"
	"log"
	"testing"
)

func TestExecutorServiceStart(t *testing.T) {
	es := executor.NewExecutionService(1, 2)
	log.Printf("es: %v\n", es)
}
