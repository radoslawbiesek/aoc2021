package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt(str string) int {
	parsed, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("Could not parse %s", str))
	}

	return parsed
}

func getInput(path string) []int {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	content := string(data)

	splittedElements := strings.Split(content, "\n")
	parsedElements := []int{}

	for _, el := range splittedElements {
		parsed := parseInt(el)
		parsedElements = append(parsedElements, parsed)
	}

	return parsedElements
}

func part1(path string) int {
	input := getInput(path)
	increased := 0

	for i := 0; i < len(input)-1; i++ {
		curr := input[i]
		next := input[i+1]

		if next > curr {
			increased++
		}
	}

	return increased
}

func part2(path string) int {
	input := getInput(path)
	increased := 0

	for i := 0; i < len(input)-3; i++ {
		curr := input[i] + input[i+1] + input[i+2]
		next := input[i+1] + input[i+2] + input[i+3]

		if next > curr {
			increased++
		}
	}

	return increased
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
