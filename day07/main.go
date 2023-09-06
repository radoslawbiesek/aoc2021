package main

import (
	"fmt"
	"math"

	"github.com/radoslawbiesek/aoc2021/utils"
)

func getInput(path string) (positions []int) {
	for _, numStr := range utils.GetLines(path, ",") {
		position := utils.ParseInt(numStr)
		positions = append(positions, position)
	}
	return
}

func calculateMinAndMaxPosition(positions []int) (min, max int) {
	for _, position := range positions {
		if position < min {
			min = position
		} else if position > max {
			max = position
		}
	}
	return
}

func part1(path string) int {
	positions := getInput(path)
	minPosition, maxPosition := calculateMinAndMaxPosition(positions)

	minCost := int(math.Inf(1))
	for targetPosition := minPosition; targetPosition <= maxPosition; targetPosition++ {
		currentCost := 0
		for _, position := range positions {
			diff := int(math.Abs(float64(position) - float64(targetPosition)))
			currentCost += diff
		}

		if currentCost < minCost {
			minCost = currentCost
		}
	}

	return minCost
}

func prepareStepsCostMap(max int) map[int]int {
	stepsCost := map[int]int{} // diff -> cost
	for diff := 1; diff <= max; diff++ {
		cost := 0
		for i := 0; i < diff; i++ {
			cost += i + 1
		}
		stepsCost[diff] = cost
	}

	return stepsCost
}

func part2(path string) int {
	positions := getInput(path)
	minPosition, maxPosition := calculateMinAndMaxPosition(positions)
	stepsCost := prepareStepsCostMap(maxPosition)

	minCost := int(math.Inf(1))
	for targetPosition := minPosition; targetPosition <= maxPosition; targetPosition++ {
		var currentCost int
		for _, position := range positions {
			diff := int(math.Abs(float64(position) - float64(targetPosition)))
			currentCost += stepsCost[diff]
		}

		if currentCost < minCost {
			minCost = currentCost
		}
	}

	return minCost
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
