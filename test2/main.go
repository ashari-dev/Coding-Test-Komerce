package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)

	fmt.Print("Input the number of families : ")
	inputFams, _ := read.ReadString(('\n'))
	inputFams = strings.TrimSpace(inputFams)
	n, err := strconv.Atoi(inputFams)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	fmt.Print("Input the number of members in the family (separated by space) : ")
	inputMembers, _ := read.ReadString('\n')
	inputMembers = strings.TrimSpace(inputMembers)
	memberStrings := strings.Split(inputMembers, " ")

	if len(memberStrings) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	family := make([]int, n)
	for i, v := range memberStrings {
		family[i], err = strconv.Atoi(v)
		if err != nil || family[i] <= 0 {
			fmt.Println("Invalid input")
			return
		}
	}

	for _, v := range family {
		if v > 4 {
			fmt.Println("families cannot be put into one bus, the members of family is more than 4")
			return
		}
	}

	result := minBuss(family)

	fmt.Println("Minimum buss : ", result)
	fmt.Println(family)
}

func minBuss(fams []int) int {
	bus := 0
	n := len(fams)

	sortSlice(fams)

	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		bus++
		visited[i] = true
		for j := i + 1; j < n; j++ {
			if !visited[j] && fams[i]+fams[j] <= 4 {
				visited[j] = true
				break
			}
		}
	}

	return bus
}

func sortSlice(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
