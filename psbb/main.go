package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	minBus(family)
}

func minBus(fams []int) {
	bus := 0
	visit := make([]bool, len(fams))
	sort.Slice(fams, func(i, j int) bool { return fams[j] < fams[i] })

	for i := 0; i < len(fams); i++ {
		if visit[i] {
			continue
		}
		bus++
		visit[i] = true
		for j := i + 1; j < len(fams); j++ {
			if !visit[j] && fams[i]+fams[j] <= 4 {
				visit[j] = true
				break
			}
		}
	}

	fmt.Println("Minimum bus required is :", bus)
}
