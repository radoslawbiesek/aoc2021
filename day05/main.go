package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

func newLine(lineStr string) utils.Line {
	pointStrings := strings.Split(lineStr, " -> ")

	startStr := pointStrings[0]
	startStrResult := strings.Split(startStr, ",")
	startX := utils.ParseInt(startStrResult[0])
	startY := utils.ParseInt(startStrResult[1])
	p1 := utils.Point{X: startX, Y: startY}

	endStr := pointStrings[1]
	endStrResult := strings.Split(endStr, ",")
	endX := utils.ParseInt(endStrResult[0])
	endY := utils.ParseInt(endStrResult[1])
	p2 := utils.Point{X: endX, Y: endY}

	return utils.Line{
		P1: p1,
		P2: p2,
	}
}

func getInput(path string) (lines []utils.Line) {
	lineStrings := utils.GetLines(path, "\n")
	for _, lineStr := range lineStrings {
		lines = append(lines, newLine(lineStr))
	}
	return
}

func countOverlappingPoints(diagram utils.Diagram, numberOfLines int) (total int) {
	for _, value := range diagram {
		if value >= numberOfLines {
			total++
		}
	}
	return
}

func part1(path string) int {
	lines := getInput(path)
	diagram := utils.Diagram{}

	for _, line := range lines {
		if !line.IsHorizontal() && !line.IsVertical() {
			continue
		}
		points := line.GetPoints()
		for _, point := range points {
			current := diagram.GetPositionValue(point.X, point.Y)
			diagram.SetPositionValue(point.X, point.Y, current+1)
		}
	}

	return countOverlappingPoints(diagram, 2)
}

func part2(path string) int {
	lines := getInput(path)
	diagram := utils.Diagram{}

	for _, line := range lines {
		points := line.GetPoints()
		for _, point := range points {
			current := diagram.GetPositionValue(point.X, point.Y)
			diagram.SetPositionValue(point.X, point.Y, current+1)
		}
	}

	return countOverlappingPoints(diagram, 2)
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
