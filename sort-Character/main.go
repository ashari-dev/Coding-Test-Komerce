package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of words (S): ")
	input, _ := read.ReadString('\n')
	input = strings.TrimSpace(input)

	sortChart(input)
}

func sortChart(str string) {
	vowel := "aiueo"
	var vowelChart, consonantChart string

	str = strings.ToLower(str)
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			if strings.ContainsRune(vowel, char) {
				vowelChart += string(char)
			} else {
				consonantChart += string(char)
			}
		}
	}

	fmt.Println("Vowel Characters : ", vowelChart)
	fmt.Println("Consonant Characters : ", consonantChart)
}
