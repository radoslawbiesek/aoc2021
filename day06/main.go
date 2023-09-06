package main

import (
	"fmt"

	"github.com/radoslawbiesek/aoc2021/utils"
)

const NEW_FISH_DAYS int = 8
const RESET_FISH_DAYS int = 6

type fishArray [NEW_FISH_DAYS + 1]int

func getInput(path string) fishArray {
	fishArr := fishArray{} // index -> day, value -> count
	lineStrings := utils.GetLines(path, ",")
	for _, numStr := range lineStrings {
		day := utils.ParseInt(numStr)
		fishArr[day]++
	}

	return fishArr
}

func calculateFishes(fishArr fishArray, days int) int {
	for i := 0; i < days; i++ {
		newFishArr := fishArray{}
		for day, count := range fishArr {
			if day == 0 {
				newFishArr[NEW_FISH_DAYS] += count
				newFishArr[RESET_FISH_DAYS] += count
			} else {
				newFishArr[day-1] += count
			}
		}
		fishArr = newFishArr
	}

	total := 0
	for _, count := range fishArr {
		total += count
	}

	return total
}

func part1(path string) int {
	fishMap := getInput(path)
	return calculateFishes(fishMap, 80)
}

func part2(path string) int {
	fishMap := getInput(path)
	return calculateFishes(fishMap, 256)
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
