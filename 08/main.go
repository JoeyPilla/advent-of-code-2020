package main

import (
	"../utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Operation string

const (
	Acc Operation = "acc"
	Jmp           = "jmp"
	Nop           = "nop"
)

type Instruction struct {
	op       Operation
	arg      int
	calledAt []int
}

type Result struct {
	acc      int
	finished bool
}

func part1(inputs []string) {
	var wg sync.WaitGroup
	results := make(chan Result)
	done := make(chan bool)
	go readResults(results, false, done)
	instructions := LoadInstructions(inputs)
	wg.Add(1)
	go ProcessInstructions(instructions, results, &wg, done)
	wg.Wait()
	close(results)
}

func part2(inputs []string) {
	var wg sync.WaitGroup
	results := make(chan Result)
	done := make(chan bool)
	go readResults(results, true, done)
	for i := 0; i < len(inputs); i++ {
		instructions := LoadInstructions(inputs)
		changed := false
		instructions[i], changed = changeInstruction(instructions[i])
		if !changed {
			continue
		}
		wg.Add(1)
		go ProcessInstructions(instructions, results, &wg, done)
	}
	wg.Wait()
	close(results)
}

func readResults(r chan Result, care bool, done chan bool) {
	for res := range r {
		if res.finished || !care {
			fmt.Println(res.acc)
			done <- true
		}
	}
}

func changeInstruction(i Instruction) (Instruction, bool) {
	switch i.op {
	case Acc:
		return i, false
	case Jmp:
		i.op = Nop
	case Nop:
		i.op = Jmp
	default:
		fmt.Errorf("invalid operation")
	}
	return i, true
}

func LoadInstructions(inputs []string) []Instruction {
	instructions := make([]Instruction, len(inputs))
	for i, input := range inputs {
		temp := strings.Split(input, " ")
		value, _ := strconv.Atoi(temp[1])
		instructions[i] = Instruction{
			op:       Operation(temp[0]),
			arg:      value,
			calledAt: []int{},
		}
	}
	return instructions
}

func ProcessInstructions(instructions []Instruction, r chan Result, wg *sync.WaitGroup, done chan bool) {
	acc, pos, count := 0, 0, 0
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			if pos > len(instructions) {
				r <- Result{
					acc:      acc,
					finished: false,
				}
				return
			}
			if pos == len(instructions) {
				r <- Result{
					acc:      acc,
					finished: true,
				}
				return
			}
			i := &instructions[pos]
			if len(i.calledAt) > 0 {
				r <- Result{
					acc:      acc,
					finished: false,
				}
				return
			}
			i.calledAt = append(i.calledAt, count)
			pos, acc = ProcessInstruction(i, pos, acc)
			count++
		}
	}
}

func ProcessInstruction(i *Instruction, pos, acc int) (int, int) {
	switch i.op {
	case Acc:
		acc += i.arg
		pos += 1
	case Jmp:
		pos += i.arg
	case Nop:
		pos += 1
	default:
		fmt.Errorf("invalid operation")
	}
	return pos, acc
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	start := time.Now()
	part1(inputs)
	elapsed := time.Since(start)
	log.Printf("Part1 took %s", elapsed)
	start = time.Now()
	part2(inputs)
	elapsed = time.Since(start)
	log.Printf("Part2 took %s", elapsed)
}
