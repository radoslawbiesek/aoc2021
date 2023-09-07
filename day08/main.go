package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

var DIGITS_MAP = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

type display struct {
	patterns, output []string
}

func getInput(path string) (displays []display) {
	lines := utils.GetLines(path, "\n")
	for _, line := range lines {
		splitted := strings.Split(line, "|")

		patterns := strings.Split(splitted[0], " ")
		patterns = patterns[:len(patterns)-1]

		output := strings.Split(splitted[1], " ")
		output = output[1:]

		displays = append(displays, display{patterns: patterns, output: output})
	}
	return
}

func part1(path string) (total int) {
	displays := getInput(path)
	for _, display := range displays {
		for _, output := range display.output {
			switch len(output) {
			case len(DIGITS_MAP[1]), len(DIGITS_MAP[4]), len(DIGITS_MAP[7]), len(DIGITS_MAP[8]):
				total++
			}
		}
	}

	return
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
}
