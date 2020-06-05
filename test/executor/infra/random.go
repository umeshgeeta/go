// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package infra

import "math/rand"

var GlobalRand *rand.Rand

func SetupRand() {
	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	GlobalRand = rand.New(rand.NewSource(99))
}

func GetRandomBoolean() bool {
	r := GlobalRand.Intn(2)
	if r == 0 {
		return false
	} else {
		return true
	}
}
