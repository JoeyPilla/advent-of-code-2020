package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputs := "17,1,3,16,19,0"
	start := time.Now()
	part1(inputs, 2020)
	elapsed := time.Since(start)
	log.Printf("Part1 took %s", elapsed)
	start = time.Now()
	part1(inputs, 30000000)
	elapsed = time.Since(start)
	log.Printf("Part1 took %s", elapsed)
}

func part1(input string, count int) {
	temp := strings.Split(input, ",")
	mapping := make(map[int]int)
	for i, num := range temp {
		number, _ := strconv.Atoi(num)
		mapping[number] = i + 1
	}
	last := 0
	for i := len(temp) + 1; i < count; i++ {
		if curr, ok := mapping[last]; ok {
			mapping[last] = i
			last = i - curr
		} else {
			mapping[last] = i
			last = 0
		}
	}
	fmt.Println(last)
}
