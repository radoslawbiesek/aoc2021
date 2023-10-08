package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

func nextStep(g *utils.Grid) (flashes int) {
	queue := utils.Queue[utils.Point]{}
	points := g.GetAllPoints()
	for _, point := range points {
		queue.Enqueue(point)
	}

	for queue.Len > 0 {
		p, _ := queue.Dequeue()
		curr := g.GetValue(*p)
		new := curr + 1
		g.SetValue(*p, new)
		if new == 10 {
			flashes++
			for _, n := range utils.Get8Neighbors(*g, *p) {
				queue.Enqueue(n)
			}
		}
	}

	for _, p := range g.GetAllPoints() {
		curr := g.GetValue(p)
		if curr >= 10 {
			g.SetValue(p, 0)
		}
	}
	return
}

func getInput(path string) (grid utils.Grid) {
	lineStrings := utils.GetLines(path, "\n")
	for _, lineStr := range lineStrings {
		line := []int{}
		chars := strings.Split(lineStr, "")
		for _, char := range chars {
			line = append(line, utils.ParseInt(char))
		}
		grid = append(grid, line)
	}
	return
}

func part1(path string) (flashes int) {
	grid := getInput(path)
	totalSteps := 100
	for step := 0; step < totalSteps; step++ {
		flashes += nextStep(&grid)
	}
	return
}

func part2(path string) (step int) {
	grid := getInput(path)
	width, height := grid.GetDimensions()
	step = 0
	for {
		step++
		flashes := nextStep(&grid)
		if flashes == width*height {
			return
		}
	}
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
