package main

import (
	"../utils"
	"fmt"
	"strconv"
)

func part1(inputs []string, buffer int) int {
	prev := make([]int, buffer)

	for count, input := range inputs {
		i, _ := strconv.Atoi(input)
		if count > buffer && !checkNumber(prev, i) {
			return i
		}
		prev[count%buffer] = i
	}
	return -1
}

func checkNumber(arr []int, num int) bool {
	m := make(map[int]bool)
	for _, i := range arr {
		if _, ok := m[i]; ok {
			return true
		}
		m[num-i] = true
	}
	return false
}

func part2(inputs []string, val int) int {
	acc, small, large := 0, 100000, 0
	for start := 0; start < len(inputs); start++ {
		for count, input := range inputs {
			if count < start {
				continue
			}
			i, _ := strconv.Atoi(input)
			acc += i

			if i < small {
				small = i
			} else if i > large {
				large = i
			}
			if acc == val {
				return small + large
			}
			if acc > val {
				small = i
				large = i
				acc = i
			}
		}
	}
	return -1
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	invalidValue := part1(inputs, 25)
	fmt.Println(invalidValue)
	fmt.Println(part2(inputs, invalidValue))

}
