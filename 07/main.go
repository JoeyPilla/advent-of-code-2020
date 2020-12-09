package main

import (
	"../utils"
	"fmt"
	"strconv"
	"strings"
)

type Bag struct {
	checked     bool
	color       string
	hasTarget   bool
	subBagCount int
}

func part1new(inputs []string) {
	m := make(map[string]Bag)
	for _, input := range inputs {
		io := strings.Split(input, "contain")
		bagType := getBagType(io[0])
		m[bagType] = processBag(bagType, io[1], m)
	}
}

func processBag(color, str string, m map[string]Bag) Bag {
	bag := Bag{
		color: color,
	}

	fmt.Println(processSubBags(str))
	return bag
}

func part1(inputs []string) {
	m := make(map[string]string)
	for _, input := range inputs {
		io := strings.Split(input, "contain")
		bagType := getBagType(io[0])
		m[bagType] = io[1]
	}
	fmt.Println(getGoldBags(m))
}

func part2(inputs []string) {
	m := make(map[string][][]string)
	for _, input := range inputs {
		io := strings.Split(input, "contain")
		bagType := getBagType(io[0])
		m[bagType] = processSubBags(io[1])
	}
	fmt.Println(getSubBagCounts(m)["shiny gold"])
}

func getSubBagCounts(m map[string][][]string) map[string]int {
	valueMap := make(map[string]int)
	for k, v := range m {
		total := 0
		for _, val := range v {
			total += countBags(m, val)
		}
		valueMap[k] = total
	}
	return valueMap
}

func countBags(m map[string][][]string, bag []string) int {
	if bag[0] == "0" {
		return 0
	}
	subBag := m[bag[1]]
	val, _ := strconv.Atoi(bag[0])
	total := val
	for _, s := range subBag {
		total += val * countBags(m, s)
	}
	return total
}

func getBagType(s string) string {
	adjectives := strings.Split(s, " ")
	return adjectives[0] + " " + adjectives[1]
}

func getBagTypeWithNum(s string) []string {
	adjectives := strings.Split(s, " ")
	return []string{adjectives[0], adjectives[1] + " " + adjectives[2]}
}

func getGoldBags(m map[string]string) int {
	total := 0
	for _, v := range m {
		subBags := processSubBags(v)
		found := false
		for _, bag := range subBags {
			if getGoldBag(m, bag[1]) {
				found = true
			}
		}
		if found {
			total += 1
		}
	}
	return total
}

func processSubBags(s string) [][]string {
	bags := strings.Split(s, ",")
	retBags := make([][]string, len(bags))
	for i, bag := range bags {
		if bag == " no other bags." {
			retBags[i] = []string{"0", "NONE"}
		} else {
			retBags[i] = getBagTypeWithNum(strings.Trim(bag, " "))
		}
	}
	return retBags
}

func getGoldBag(m map[string]string, bag string) bool {
	if bag == "NONE" {
		return false
	} else if bag == "shiny gold" {
		return true
	} else {
		found := false
		bags := processSubBags(m[bag])
		for _, subBag := range bags {
			found = found || getGoldBag(m, subBag[1])
		}
		return found
	}
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	part1new(inputs)
	part2(inputs)
}
