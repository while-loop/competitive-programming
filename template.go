package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 *
 */
func main() {
	in.Split(bufio.ScanWords)

	N := nextInt()
	for i := 0; i < N; i++ {

	}
	fmt.Println(N)
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}

func nextFloat64() float64 {
	in.Scan()
	tmp, _ := strconv.ParseFloat(in.Text(), 64)
	return tmp
}

func nextLine() string {
	in.Scan()
	return in.Text()
}

func next() string {
	in.Scan()
	return in.Text()
}
