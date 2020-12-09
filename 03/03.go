package main

import (
	"../utils"
	"fmt"
)

func part1(inputs []string, right, down int) int {
	length := len(inputs[0])
	x, y, count := 0, 0, 0

	for y < len(inputs) {
		input := inputs[y]
		if '#' == input[x%length] {
			count += 1
		}
		x += right
		y += down
	}
	return count
}

func part2(inputs []string) int {
	answer := 1
	runs := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, run := range runs {
		answer *= part1(inputs, run[0], run[1])
	}
	return answer
}

func main() {
	inputs := utils.ReadTextFile("./03-input.txt")
	fmt.Println(part1(inputs, 3, 1))
	fmt.Println(part2(inputs))
}
