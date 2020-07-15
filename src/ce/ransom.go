// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import "fmt"

func main() {
	rn := "a"
	mag := "b"
	fmt.Printf("ransome note: %s magazine: %s result: %v\n", rn, mag, ransomeNode(rn, mag))

	rn = "aa"
	mag = "ab"
	fmt.Printf("ransome note: %s magazine: %s result: %v\n", rn, mag, ransomeNode(rn, mag))

	rn = "abab"
	mag = "baab"
	fmt.Printf("ransome note: %s magazine: %s result: %v\n", rn, mag, ransomeNode(rn, mag))
}

func ransomeNode(rn string, mag string) bool {
	if rn == "" {
		return true
	}
	if mag == "" {
		return false
	}
	if len(mag) < len(rn) {
		// we have insufficient characters in mag string
		return false
	}
	magLetters := make(map[rune]int)
	for _, v := range mag {
		c, ok := magLetters[v]
		if !ok {
			magLetters[v] = 1
		} else {
			magLetters[v] = c + 1
		}
	}
	for _, c := range rn {
		cc, ok := magLetters[c]
		if !ok {
			return false
		} else {
			magLetters[c] = cc - 1
			if magLetters[c] == 0 {
				delete(magLetters, c)
			}
		}
	}
	return true
}
