package main

import (
	"../utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part1(inputs []string) {
	m := make(map[string]int)
	total := 0
	count := 0
	for _, input := range inputs {
		if input == "" {
			for _, val := range m {
				if val == count {
					total += 1
				}
			}
			m = make(map[string]int)
			count = 0
		} else {
			for _, s := range input {
				val, ok := m[string(s)]
				if !ok {
					test[string(s)] = 1
				} else {
					test[string(s)] = val + 1
				}
			}
			count += 1
		}
	}
	fmt.Println(total)
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	part1(inputs)
}
