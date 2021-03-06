package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/mark-the-answer-1/
 */
func main1() {
	fmt.Print(SolveMark(bufio.NewScanner(os.Stdin)))
}

func SolveMark(scanner *bufio.Scanner) int {
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	tests, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	max, _ := strconv.Atoi(scanner.Text())
	skipped, count := false, 0
	var tmp int
	for i := 0; i < tests; i++ {
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())

		if tmp <= max {
			count++
		} else if !skipped {
			skipped = true
		} else {
			break
		}
	}
	return count
}
