package main

import (
	"../utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part1(inputs []string) int {
	num, _ := strconv.Atoi(inputs[0])
	min := 5432543544
	bus, wait := 0, 0
	arr := strings.Split(inputs[1], ",")
	for _, val := range arr {
		if val == "x" {
			continue
		}

		num2, _ := strconv.Atoi(val)
		if num+num2-num%num2 < min {
			bus = num2
			min = num + num2 - num%num2
			wait = bus - num%num2
		}
	}

	fmt.Println(bus * wait)
	return num
}

func crt(pairs map[int]int) int {
	M := 1
	total := 0
	for x, mx := range pairs {
		fmt.Println(x, mx)
		M *= mx
		total = 0
	}
	for x, mx := range pairs {
		b := M / mx
		total += x * b * int(math.Pow(float64(b), float64(mx-2))) % mx
		total %= M
	}
	return total
}

func part2(inputs []string) int {

	arr := strings.Split(inputs[1], ",")
	pairs := make(map[int]int)
	for i, val := range arr {
		if val == "x" {
			continue
		}

		n, _ := strconv.Atoi(val)
		pairs[n-i] = n
	}
	fmt.Println(pairs)
	fmt.Println(crt(pairs))
	return crt(pairs)
}

func pafdsart2(inputs []string) int {
	num, _ := strconv.Atoi(inputs[0])
	offsets := make(map[int]int)
	numbers := []int{}
	arr := strings.Split(inputs[1], ",")
	for i, val := range arr {
		if val == "x" {
			continue
		}
		num2, _ := strconv.Atoi(val)
		numbers = append(numbers, num2)
		if i == 0 {
			offsets[num2] = num2
		} else {
			offsets[num2] = i
		}
	}
	i := 0
	fmt.Println(offsets)
	for {
		finished := true
		for _, n := range numbers {
			if i%n != n-offsets[n] {
				finished = false
				break
			}
		}
		if finished {
			break
		}
		i++
	}

	fmt.Println(i)
	return num
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	part1(inputs)
	part2(inputs)
}
