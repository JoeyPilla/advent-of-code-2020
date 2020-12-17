package main

import (
	"../utils"
	"fmt"
	"math"
	"strconv"
	//	"strings"
)

type Instruction struct {
	direction string
	value     int
}

func part1(inputs []string) {
	instructions := processInput(inputs)
	x, y := processInstructions(instructions)
	fmt.Println(x, y)
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func processInput(inputs []string) []Instruction {
	instructions := make([]Instruction, len(inputs))
	for i, input := range inputs {
		dir := string(input[0])
		value, _ := strconv.Atoi(string(input[1:]))
		instructions[i] = Instruction{
			direction: dir,
			value:     value,
		}
	}
	return instructions
}

func makePositive(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func makeNegative(n int) int {
	if n > 0 {
		return -n
	}
	return n
}

func processDir(dir, wayX, wayY int) (int, int) {
	directions := []string{"N", "E", "S", "W"}
	if directions[dir] == "N" {
		wayY = makePositive(wayY)
	} else if directions[dir] == "S" {
		wayY = makeNegative(wayY)
	} else if directions[dir] == "E" {
		wayX = makePositive(wayX)
	} else if directions[dir] == "W" {
		wayX = makeNegative(wayX)
	}
	fmt.Println(directions[dir], wayX, wayY)
	return wayX, wayY
}

func processInstructions(instructions []Instruction) (int, int) {
	x, y := 0, 0
	wayX, wayY := 10, 1
	dir1, dir2 := 1, 0
	for _, instr := range instructions {
		if instr.direction == "R" {
			amount := instr.value / 90
			if !(amount%2 == 0) {
				wayX, wayY = wayY, wayX
			}
			dir1 += amount
			dir2 += amount
			dir1 %= 4
			dir2 %= 4
			wayX, wayY = processDir(dir1, wayX, wayY)
			wayX, wayY = processDir(dir2, wayX, wayY)
		}
		if instr.direction == "L" {
			amount := instr.value / 90
			dir1 -= amount
			dir2 -= amount
			if dir1 < 0 {
				dir1 += 4
			}
			if dir2 < 0 {
				dir2 += 4
			}
			wayX, wayY = processDir(dir1, wayX, wayY)
			wayX, wayY = processDir(dir2, wayX, wayY)
		}

		if instr.direction == "F" {
			x += wayX * instr.value
			y += wayY * instr.value
		}

		switch instr.direction {
		case "N":
			wayY += instr.value
		case "S":
			wayY -= instr.value
		case "E":
			wayX += instr.value
		case "W":
			wayX -= instr.value
		}
	}
	return x, y
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	part1(inputs)
}
