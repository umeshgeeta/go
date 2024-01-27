// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package ace

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SubstringSearch searches for the given pattern in the given file.
// It returns a slice of integers where each element represents the
// number of occurrences of the pattern in the corresponding line of
// the file. It also returns a slice of strings where each element
// represents the ending prefix of the line that is combined with the
// beginning of the next line to form the pattern.
func SubstringSearch(fileName string, pattern string) ([]int, []string) {
	var result []int
	var endingPrefix []string

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return result, endingPrefix
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var prevLine string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, 0)
			endingPrefix = append(endingPrefix, "")
			prevLine = ""
			continue
		}
		count := strings.Count(line, pattern)
		result = append(result, count)
		if prevLine != "" {
			prevLineContribution := prevLine
			if len(prevLine) > len(pattern)-1 {
				prevLineContribution = prevLine[len(prevLine)-len(pattern)+1:]

			}
			plcLength := len(prevLineContribution)
			contribution := line
			if len(line) > len(pattern)-1 {
				contribution = line[:len(pattern)-1]
			}
			combinedLine := prevLineContribution + contribution
			if strings.Contains(combinedLine, pattern) {
				patternStart := strings.Index(combinedLine, pattern)
				suffix := pattern[plcLength-patternStart:]
				endingPrefix = append(endingPrefix, suffix)
			} else {
				endingPrefix = append(endingPrefix, "")
			}
		}

		prevLine = line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	return result, endingPrefix
}
