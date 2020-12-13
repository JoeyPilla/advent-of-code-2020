package main

import (
	"../utils"
	"fmt"
	"strings"
)

func part1(inputs []string) {
	grid := createGrid(inputs)
	processGrid(grid)
	countGrid(grid)
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func countGrid(grid [][]string) {
	count := 0
	for _, row := range grid {
		for _, c := range row {
			if c == "#" {
				count++
			}
		}

	}
	fmt.Println(count)
}

func createGrid(in []string) [][]string {
	resp := make([][]string, len(in))
	for i, row := range in {
		resp[i] = strings.Split(row, "")
	}
	return resp
}

func copyGrid(grid [][]string, newGrid [][]string) {
	for i := range grid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
}

func processGrid(grid [][]string) [][]string {
	changed := true
	for changed {
		newGrid := make([][]string, len(grid))
		copyGrid(grid, newGrid)
		changed = false
		for i, r := range grid {
			for j, c := range r {
				if c == "L" {
					if checkEmpty(grid, i, j) {
						newGrid[i][j] = "#"
						changed = true
					}
				} else if c == "#" {
					if checkFilled(grid, i, j) {
						newGrid[i][j] = "L"
						changed = true
					}
				}
			}
		}

		copyGrid(newGrid, grid)
	}
	return grid
}

func checkEmpty(grid [][]string, r, c int) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			adjSeat := checkSide(grid, i+r, j+c)
			if adjSeat == "#" {
				return false
			}
		}
	}
	return true
}

func checkSide(grid [][]string, r, c int) string {
	if r < len(grid) && r >= 0 {
		row := grid[r]
		if c < len(row) && c >= 0 {
			return row[c]
		}
	}
	return ""
}

func checkFilled(grid [][]string, r, c int) bool {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			adjSeat := checkSide(grid, i+r, j+c)
			if adjSeat == "#" {
				count++
				if count >= 4 {
					return true
				}
			}
		}
	}
	return false
}

func part2(inputs []string) {
	grid := createGrid(inputs)
	processGrid2(grid)
	countGrid(grid)
}

func processGrid2(grid [][]string) [][]string {
	changed := true
	for changed {
		newGrid := make([][]string, len(grid))
		copyGrid(grid, newGrid)
		changed = false
		for i, r := range grid {
			for j, c := range r {
				if c == "L" {
					if checkEmpty2(grid, i, j) {
						newGrid[i][j] = "#"
						changed = true
					}
				} else if c == "#" {
					if checkFilled2(grid, i, j) {
						newGrid[i][j] = "L"
						changed = true
					}
				}
			}
		}

		copyGrid(newGrid, grid)
	}
	return grid
}

func checkEmpty2(grid [][]string, r, c int) bool {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			adjSeat := checkDirection(grid, i, j, r, c)
			if adjSeat == "#" {
				return false
			}
		}
	}
	return true
}

func checkDirection(grid [][]string, i, j, r, c int) string {
	iCount, jCount := i+r, j+c
	for {
		seat := checkSide(grid, iCount, jCount)
		if seat == "" {
			return ""
		} else if seat == "L" {
			return "L"
		} else if seat == "#" {
			return "#"
		}
		iCount += i
		jCount += j
	}
	return ""
}

func checkSide2(grid [][]string, r, c int) string {
	if r < len(grid) && r >= 0 {
		row := grid[r]
		if c < len(row) && c >= 0 {
			return row[c]
		}
	}
	return ""
}

func checkFilled2(grid [][]string, r, c int) bool {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			adjSeat := checkDirection(grid, i, j, r, c)
			if adjSeat == "#" {
				count++
				if count >= 5 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	part1(inputs)
	part2(inputs)
}
