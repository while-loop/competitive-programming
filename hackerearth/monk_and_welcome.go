package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-welcome-problem/
  */
func main() {
	reader := bufio.NewReader(os.Stdin)

	tmp, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(tmp))

	tmp, _ = reader.ReadString('\n')
	a := strings.Split(tmp, " ")
	tmp, _ = reader.ReadString('\n')
	b := strings.Split(tmp, " ")

	for i := 0; i < n; i++ {
		l, _ := strconv.Atoi(strings.TrimSpace(a[i]))
		r, _ := strconv.Atoi(strings.TrimSpace(b[i]))
		fmt.Print(l + r, " ")
	}
}