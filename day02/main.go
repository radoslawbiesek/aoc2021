package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

type instruction struct {
	direction direction
	amount    int
}

type direction string

const (
	forward direction = "forward"
	up      direction = "up"
	down    direction = "down"
)

func parseInstruction(str string) instruction {
	splittedEl := strings.Split(str, " ")

	var direction direction
	switch splittedEl[0] {
	case "forward":
		direction = forward
	case "up":
		direction = up
	case "down":
		direction = down
	default:
		panic(fmt.Sprintf("Invalid direction %s", direction))
	}
	amount := utils.ParseInt(splittedEl[1])

	return instruction{direction: direction, amount: amount}
}

func getInput(path string) []instruction {
	lines := utils.GetLines(path)
	parsedElements := utils.Map(lines, parseInstruction)

	return parsedElements
}

func part1(path string) int {
	instructions := getInput(path)

	depth := 0
	horizontal := 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case forward:
			horizontal += instruction.amount
		case down:
			depth += instruction.amount
		case up:
			depth -= instruction.amount
		}
	}

	return depth * horizontal
}

func part2(path string) int {
	instructions := getInput(path)

	depth := 0
	horizontal := 0
	aim := 0

	for _, instruction := range instructions {
		switch instruction.direction {
		case forward:
			horizontal += instruction.amount
			depth += aim * instruction.amount
		case down:
			aim += instruction.amount
		case up:
			aim -= instruction.amount
		}
	}

	return depth * horizontal
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./test-input.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
	fmt.Printf("Part 2: %d\n", part2("./input.txt"))
}
