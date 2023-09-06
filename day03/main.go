package main

import (
	"fmt"
	"math"

	"github.com/radoslawbiesek/aoc2021/utils"
)

func getInput(path string) (lines []string) {
	lines = utils.GetLines(path, "\n")
	return
}

func countZerosAndOnes(lines []string, colIndex int) (zeros int, ones int) {
	for _, line := range lines {
		char := utils.CharAt(line, colIndex)
		if char == "0" {
			zeros++
		} else {
			ones++
		}
	}
	return
}

func binaryToDec(binary string) (dec int) {
	binaryLen := len(binary)
	for colIndex := 0; colIndex < binaryLen; colIndex++ {
		p := binaryLen - colIndex - 1
		currDec := int(math.Pow(2, float64(p)))
		char := utils.CharAt(binary, colIndex)
		if char == "1" {
			dec += currDec
		}
	}
	return
}

func part1(path string) int {
	lines := getInput(path)
	linesLen := len(lines[0])

	gamma := ""
	epsilon := ""

	for colIndex := 0; colIndex < linesLen; colIndex++ {
		zeros, ones := countZerosAndOnes(lines, colIndex)
		if zeros > ones {
			epsilon += "0"
			gamma += "1"
		} else {
			epsilon += "1"
			gamma += "0"
		}
	}

	gammaDec := binaryToDec(gamma)
	epsilonDec := binaryToDec(epsilon)

	return gammaDec * epsilonDec
}

func part2(path string) int {
	lines := getInput(path)
	lineLen := len(lines[0])

	oxygenRating := lines
	for len(oxygenRating) > 1 {
		for colIndex := 0; colIndex < lineLen && len(oxygenRating) > 1; colIndex++ {
			zeros, ones := countZerosAndOnes(oxygenRating, colIndex)
			oxygenRating = utils.Filter(oxygenRating, func(line string) bool {
				char := utils.CharAt(line, colIndex)
				if ones >= zeros {
					return char == "1"
				} else {
					return char == "0"
				}
			})
		}
	}

	co2ScrubberRating := lines
	for len(co2ScrubberRating) > 1 {
		for colIndex := 0; colIndex < lineLen && len(co2ScrubberRating) > 1; colIndex++ {
			zeros, ones := countZerosAndOnes(co2ScrubberRating, colIndex)
			co2ScrubberRating = utils.Filter(co2ScrubberRating, func(line string) bool {
				char := utils.CharAt(line, colIndex)
				if zeros <= ones {
					return char == "0"
				} else {
					return char == "1"
				}
			})
		}
	}

	oxygenRatingDec := binaryToDec(oxygenRating[0])
	co2ScrubberRatingDec := binaryToDec(co2ScrubberRating[0])

	return oxygenRatingDec * co2ScrubberRatingDec
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
