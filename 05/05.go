package main

import (
	"../utils"
	"fmt"
	"sort"
	"strconv"
)

func part1(inputs []string) {
	max := 0
	for _, input := range inputs {
		id := getID(input[:7], input[7:])
		if id > max {
			max = id
		}
	}
	fmt.Println(max)
}

func part2(inputs []string) int {
	arr := make([]int, len(inputs))
	for i, input := range inputs {
		id := getID(input[:7], input[7:])
		arr[i] = id
	}
	sort.Ints(arr)
	for i, val := range arr {
		if i+32 != val {
			return i + 32
		}
	}
	return -1
}

func getID(row, col string) int {
	return getRow(row)*8 + getCol(col)
}

func getRow(row string) int {
	start := ""
	for _, char := range row {
		if string(char) == "B" {
			start += "1"
		} else {
			start += "0"
		}
	}
	return biStrToInt(start)
}

func getCol(col string) int {
	start := ""
	for _, char := range col {
		if string(char) == "R" {
			start += "1"
		} else {
			start += "0"
		}
	}
	return biStrToInt(start)
}

func biStrToInt(str string) int {
	if i, err := strconv.ParseInt(str, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		return int(i)
	}
	return 0
}

func main() {
	inputs := utils.ReadTextFile("./05-input.txt")
	part1(inputs)
	fmt.Println(part2(inputs))
}
