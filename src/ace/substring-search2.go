package ace

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func SubstringSearch2(fileame, word string) []int {
	file, err := os.Open(fileame)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]int, 0)
	lineNo := 0
	var wg sync.WaitGroup

	for scanner.Scan() {
		result = append(result, 0)
		line := scanner.Text()
		wg.Add(1)
		go func(l string, ln int) {
			defer wg.Done()
			count := strings.Count(l, word)
			result[ln] = count
		}(line, lineNo)
		lineNo++
	}
	wg.Wait()
	return result
}
