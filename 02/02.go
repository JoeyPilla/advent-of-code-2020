package main

import (
	"../utils"
	"fmt"
	"strconv"
	"strings"
)

func part1(inputs []string) {
	number := 0
	for _, input := range inputs {

		splitInput := strings.Split(input, " ")
		criteria, letter, str := splitInput[0], splitInput[1][0], splitInput[2]
		count := strings.Count(str, string(letter))
		strA := strings.Split(criteria, "-")
		lower, _ := strconv.Atoi(strA[0])
		upper, _ := strconv.Atoi(strA[1])
		if count >= lower && count <= upper {
			number += 1
		}
	}
	fmt.Println(number)
}

func part2(inputs []string) {
	number := 0
	for _, input := range inputs {

		splitInput := strings.Split(input, " ")
		criteria, letter, str := splitInput[0], splitInput[1][0], splitInput[2]
		strA := strings.Split(criteria, "-")
		lower, _ := strconv.Atoi(strA[0])
		upper, _ := strconv.Atoi(strA[1])
		lowerChar := str[lower-1]
		upperChar := str[upper-1]
		if lowerChar == letter && upperChar != letter || lowerChar != letter && upperChar == letter {
			number += 1
		}
	}
	fmt.Println(number)
}

func main() {
	inputs := utils.ReadTextFile("./02-input.txt")
	part1(inputs)
	part2(inputs)
}
