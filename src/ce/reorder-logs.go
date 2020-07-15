// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"strings"
)

//
// You have an array of logs.  Each log is a space delimited string of words.
//
// For each log, the first word in each log is an alphanumeric identifier.
// Then, either:
//
// Each word after the identifier will consist only of lowercase letters, or;
// Each word after the identifier will consist only of digits.
//
// We will call these two varieties of logs letter-logs and digit-logs.
// It is guaranteed that each log has at least one word after its identifier.
//
// Reorder the logs so that all of the letter-logs come before any digit-log.
// The letter-logs are ordered lexicographically ignoring identifier, with the
// identifier used in case of ties.  The digit-logs should be put in their
// original order.
//
// Return the final order of the logs.
//
// Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
// Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
//
// https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/2974/
//

const zeroRune rune = '0'
const nineRune rune = '9'

type log struct {
	identifier   string
	message      string
	isDigitalLog bool
}

func createLog(s string) *log {
	result := new(log)
	words := strings.Fields(s)
	result.identifier = words[0]
	ss := strings.SplitN(s, " ", 2)
	result.message = ss[1]
	mr := []rune(result.message)
	if zeroRune <= mr[0] && mr[0] <= nineRune {
		result.isDigitalLog = true
	}
	// else it is false i.e. letter log
	return result
}

// True if l is smaller than l2
func (l *log) isSmaller(l2 *log) bool {
	result := false
	c := strings.Compare(l.message, l2.message)
	switch c {
	case -1:
		result = true
	case 1: // l is bigger than l2, we keep result as is false
	case 0: // both messages are equal, we need to compare on id
		ci := strings.Compare(l.identifier, l2.identifier)
		switch ci {
		case -1:
			result = true
			// for id equal or id of l greater than l2; we keep it false
		}
	}
	return result
}

func (l *log) toString() string {
	return l.identifier + " " + l.message
}
func reorderLogFiles(logs []string) []string {
	logArray := make([]log, len(logs))
	digitLogStart := 0
	digitLogEnd := 0
	for _, ls := range logs {
		logElement := createLog(ls)
		if logElement.isDigitalLog {
			logArray[digitLogEnd] = *logElement
		} else {
			j := 0
			done := false
			for j < digitLogStart && !done {
				if logArray[j].isSmaller(logElement) {
					j++
				} else {
					// we insert here
					copy(logArray[j+1:], logArray[j:])
					logArray[j] = *logElement
					done = true
				}
			}
			if !done {
				// we traversed all letter logs and the incoming log is greater than
				// all of those and we have hit digitLogStart; we insert there
				copy(logArray[digitLogStart+1:], logArray[digitLogStart:])
				logArray[digitLogStart] = *logElement
			}
			// we got the letter log and we inserted it, so start of digiLog must be moved
			digitLogStart++
		}
		// regardless of whether digital log or letter log; index for the digitEnd moves
		digitLogEnd++
	}
	result := make([]string, len(logs))
	for i, le := range logArray {
		result[i] = le.toString()
	}
	return result
}

func main() {
	input := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	result := reorderLogFiles(input)
	fmt.Printf("%v\n", result)
}
