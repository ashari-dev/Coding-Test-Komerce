package library

import (
	"fmt"
	"sort"
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

	chartVowwl := strings.Split(vowelChart, "")
	sort.Strings(chartVowwl)

	fmt.Println("Vowel Characters : ", strings.Join(chartVowwl, ""))
	fmt.Println("Consonant Characters : ", consonantChart)
}
