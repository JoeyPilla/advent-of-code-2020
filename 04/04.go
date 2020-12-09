package main

import (
	"../utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part1(inputs []string) {
	test := make(map[string]string)
	total := 0
	for _, input := range inputs {
		if input == "" {
			total += checkTest(test)
			test = make(map[string]string)
		} else {
			for _, s := range strings.Split(input, " ") {
				key, val := checkPair(s)
				test[key] = val
			}
		}
	}
	fmt.Println(total)
}

func checkPair(str string) (string, string) {
	s := strings.Split(str, ":")
	return s[0], s[1]
}

func checkTest(m map[string]string) int {
	arr := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, key := range arr {
		val, ok := m[key]
		if !ok {
			return 0
		}
		if !checkValue(key, val) {
			return 0
		}
	}
	return 1
}

func checkValue(key, value string) bool {
	switch key {
	case "byr":
		return checkByr(value)
	case "iyr":
		return checkIyr(value)
	case "eyr":
		return checkEyr(value)
	case "hgt":
		return checkHgt(value)
	case "hcl":
		return checkHcl(value)
	case "ecl":
		return checkEcl(value)
	case "pid":
		return checkPid(value)
	}
	return false
}

func checkByr(val string) bool {
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if i < 1920 || i > 2002 {
		return false
	}
	return true
}

func checkIyr(val string) bool {
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if i < 2010 || i > 2020 {
		return false
	}
	return true
}

func checkEyr(val string) bool {
	i, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if i < 2020 || i > 2030 {
		return false
	}
	return true
}

func checkHgt(val string) bool {
	matched, _ := regexp.MatchString(`^[1][5-8][0-9][c][m]|[1][9][0-3][c][m]|[5][9][i][n]|[6][0-9][i][n]|[7][0-6][i][n]`, val)
	return matched
}

func checkHcl(val string) bool {
	matched, _ := regexp.MatchString(`^[#][a-f0-9][a-f0-9][a-f0-9][a-f0-9][a-f0-9][a-f0-9]`, val)
	return matched
}
func checkEcl(val string) bool {
	matched, _ := regexp.MatchString(`\bamb|blu|brn|gry|grn|hzl|oth\b`, val)
	return matched
}
func checkPid(val string) bool {
	matched, _ := regexp.MatchString(`^\d{9}$`, val)
	return matched
}

func main() {
	inputs := utils.ReadTextFile("./04-input.txt")
	part1(inputs)
}
