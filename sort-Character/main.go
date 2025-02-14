package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"test1/library"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of words (S): ")
	input, _ := read.ReadString('\n')
	input = strings.TrimSpace(input)

	library.SortChart(input)
}
