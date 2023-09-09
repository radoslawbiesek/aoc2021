package main

import (
	"fmt"
	"sort"
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
	patterns, outputs []string
}

func sortChars(str string) string {
	var slice sort.StringSlice = strings.Split(str, "")
	slice.Sort()

	return strings.Join(slice, "")
}

func getInput(path string) (displays []display) {
	lines := utils.GetLines(path, "\n")
	for _, line := range lines {
		splitted := strings.Split(line, "|")
		patterns := strings.Split(strings.Trim(splitted[0], " "), " ")
		patterns = utils.Map(patterns, sortChars)
		outputs := strings.Split(strings.Trim(splitted[1], " "), " ")
		outputs = utils.Map(outputs, sortChars)

		displays = append(displays, display{patterns: patterns, outputs: outputs})
	}
	return
}

func part1(path string) (total int) {
	displays := getInput(path)
	for _, display := range displays {
		for _, output := range display.outputs {
			switch len(output) {
			case len(DIGITS_MAP[1]), len(DIGITS_MAP[4]), len(DIGITS_MAP[7]), len(DIGITS_MAP[8]):
				total++
			}
		}
	}

	return
}

func getPatternDiff(baseStr, comparedStr string) (diff string) {
	for _, char := range strings.Split(comparedStr, "") {
		if !strings.Contains(baseStr, char) {
			diff += char
		}
	}
	return
}

func deleteIndex(slice []string, index int) []string {
	lastIndex := len(slice) - 1
	slice[index] = slice[lastIndex]
	slice = slice[:lastIndex]
	return slice
}

func checkNum(m map[int]string, num int) {
	_, ok := m[num]
	if !ok {
		msg := fmt.Sprintf("Could not find number: %v", num)
		panic(msg)
	}
}

func part2(path string) (total int) {
	displays := getInput(path)

	for _, display := range displays {
		digitsMap := map[int]string{}
		remainingPatterns := []string{}

		// find 1, 4, 7, 8
		for _, pattern := range display.patterns {
			switch len(pattern) {
			case len(DIGITS_MAP[1]):
				digitsMap[1] = pattern
			case len(DIGITS_MAP[4]):
				digitsMap[4] = pattern
			case len(DIGITS_MAP[7]):
				digitsMap[7] = pattern
			case len(DIGITS_MAP[8]):
				digitsMap[8] = pattern
			default:
				remainingPatterns = append(remainingPatterns, pattern)
			}
		}

		// find 6 based on 1
		for index, pattern := range remainingPatterns {
			diff := getPatternDiff(pattern, digitsMap[1])
			if len(diff) == 1 && len(pattern) == len(DIGITS_MAP[6]) {
				digitsMap[6] = pattern
				remainingPatterns = deleteIndex(remainingPatterns, index)
				break
			}
		}
		checkNum(digitsMap, 6)

		// find 5 based on 6
		for index, pattern := range remainingPatterns {
			diff := getPatternDiff(pattern, digitsMap[6])
			if len(diff) == 1 && len(pattern) == len(DIGITS_MAP[5]) {
				digitsMap[5] = pattern
				remainingPatterns = deleteIndex(remainingPatterns, index)
				break
			}
		}
		checkNum(digitsMap, 5)

		// find 9 based on 5
		for index, pattern := range remainingPatterns {
			diff := getPatternDiff(digitsMap[5], pattern)
			if len(diff) == 1 && len(pattern) == len(DIGITS_MAP[9]) {
				digitsMap[9] = pattern
				remainingPatterns = deleteIndex(remainingPatterns, index)
				break
			}
		}
		checkNum(digitsMap, 9)

		// find 0 based on 4
		for index, pattern := range remainingPatterns {
			diff := getPatternDiff(pattern, digitsMap[4])
			if len(diff) == 1 && len(pattern) == len(DIGITS_MAP[0]) {
				digitsMap[0] = pattern
				remainingPatterns = deleteIndex(remainingPatterns, index)
				break
			}
		}
		checkNum(digitsMap, 0)

		// find 3 based on 9
		for index, pattern := range remainingPatterns {
			diff := getPatternDiff(digitsMap[9], pattern)
			diff2 := getPatternDiff(pattern, digitsMap[9])
			if len(diff) == 0 && len(diff2) == 1 && len(pattern) == len(DIGITS_MAP[3]) {
				digitsMap[3] = pattern
				remainingPatterns = deleteIndex(remainingPatterns, index)
				break
			}
		}
		checkNum(digitsMap, 3)

		// find 2
		digitsMap[2] = remainingPatterns[0]

		patternsMap := map[string]int{}
		for digit, pattern := range digitsMap {
			patternsMap[pattern] = digit
		}

		outputNumStr := ""
		for _, output := range display.outputs {
			val, ok := patternsMap[output]
			if !ok {
				msg := fmt.Sprintf("Unknown pattern: %v", output)
				panic(msg)
			}
			outputNumStr += fmt.Sprint(val)
		}
		total += utils.ParseInt(outputNumStr)
	}

	return
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
