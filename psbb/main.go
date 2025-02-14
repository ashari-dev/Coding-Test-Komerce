package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"test2/library"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Input the number of families : ")
	num, _ := scanner.ReadString('\n')
	num = strings.TrimSpace(num)
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	fmt.Print("Input the number of members in the family (separated by a space) : ")
	numMembers, _ := scanner.ReadString('\n')
	numMembers = strings.TrimSpace(numMembers)
	members := strings.Split(numMembers, " ")

	if len(members) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	family := make([]int, n)

	for i, v := range members {
		family[i], err = strconv.Atoi(v)
		if err != nil || family[i] <= 0 {
			fmt.Println("Invalid Input")
			return
		}
		if family[i] > 4 {
			fmt.Println("Each family member must be less than or equal to 4")
			return
		}
	}

	library.MinBus(family)
}
