package library

import (
	"fmt"
	"strings"
)

func SortChart(str string) {
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
