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
	err      bool
}

type Result struct {
	acc      int
	finished bool
}

type Program struct {
	acc          int
	done         chan bool
	instrCount   int
	instruction  *Instruction
	instructions []Instruction
	kill         chan bool
	pc           int
	results      chan Result
	wg           *sync.WaitGroup
}

func CreateProgram(instructions []Instruction, results chan Result, kill chan bool, wg *sync.WaitGroup) Program {
	return Program{
		acc:          0,
		done:         make(chan bool),
		instrCount:   0,
		instruction:  &Instruction{},
		instructions: instructions,
		kill:         kill,
		pc:           0,
		results:      results,
		wg:           wg,
	}
}

func (p *Program) Run() {
	defer p.wg.Done()
	for {
		select {
		case <-p.kill:
			return
		case <-p.done:
			return
		default:
			p.readInstruction()
		}
	}
}

func (p *Program) getInstruction() *Instruction {
	if p.pc > len(p.instructions) {
		p.results <- Result{
			acc:      p.acc,
			finished: false,
		}
		return &Instruction{
			err: true,
		}
	}
	if p.pc == len(p.instructions) {
		p.results <- Result{
			acc:      p.acc,
			finished: true,
		}
		return &Instruction{
			err: true,
		}
	}
	return &p.instructions[p.pc]
}

func (p *Program) readInstruction() {
	p.instruction = p.getInstruction()
	if p.instruction.err || len(p.instruction.calledAt) > 0 {
		p.results <- Result{
			acc:      p.acc,
			finished: false,
		}
		close(p.done)
		return
	}
	p.instruction.calledAt = append(p.instruction.calledAt, p.instrCount)
	p.processInstruction()
	p.instrCount += 1
}

func (p *Program) processInstruction() {
	i := p.instruction
	switch i.op {
	case Acc:
		p.acc += i.arg
		p.pc += 1
	case Jmp:
		p.pc += i.arg
	case Nop:
		p.pc += 1
	default:
		fmt.Errorf("invalid operation")
	}
}

func part1(inputs []string) {
	var wg sync.WaitGroup
	results := make(chan Result)
	kill := make(chan bool)
	go readResults(results, false, kill)
	instructions := LoadInstructions(inputs)
	wg.Add(1)
	p := CreateProgram(instructions, results, kill, &wg)
	go p.Run()
	wg.Wait()
	close(results)
}

func part2(inputs []string) {
	var wg sync.WaitGroup
	results := make(chan Result, len(inputs))
	kill := make(chan bool)
	go readResults(results, true, kill)
	for i := 0; i < len(inputs); i++ {
		instructions := LoadInstructions(inputs)
		changed := false
		instructions[i], changed = changeInstruction(instructions[i])
		if !changed {
			continue
		}
		wg.Add(1)
		p := CreateProgram(instructions, results, kill, &wg)
		go p.Run()
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
