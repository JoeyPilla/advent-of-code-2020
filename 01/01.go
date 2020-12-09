package main

import (
	"../utils"
	"fmt"
	"strconv"
)

func find2Nums(inputs []string, val int) int {
	matches := make(map[int]bool)
	for _, strNum := range inputs {
		num, _ := strconv.Atoi(strNum)
		comp := val - num
		if _, ok := matches[num]; ok {
			return num * comp
		} else {
			matches[comp] = true
		}
	}
	return -1
}

func find3Nums(inputs []string, val int) int {
	for _, strNum := range inputs {
		num, _ := strconv.Atoi(strNum)
		comp := val - num
		if find2Nums(inputs, comp) > -1 {
			return find2Nums(inputs, comp) * num
		}
	}
	return 0
}

func main() {
	inputs := utils.ReadTextFile("../01/01-input.txt")
	fmt.Println(find2Nums(inputs, 2020))
	fmt.Println(find3Nums(inputs, 2020))
}
