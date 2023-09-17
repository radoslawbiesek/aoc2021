package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

type point struct {
	x, y int
}

type direction struct {
	x, y int
}

var directions = [8]direction{
	{y: -1, x: -1}, // NW
	{y: -1, x: 0},  // N
	{y: -1, x: 1},  // NE
	{y: 0, x: 1},   // E
	{y: 0, x: -1},  // W
	{y: 1, x: 0},   // S
	{y: 1, x: -1},  // SW
	{y: 1, x: 1},   // SE
}

type grid [10][10]int

func (g *grid) getDimensions() (width, height int) {
	width = len(g[0])
	height = len(g)
	return
}

func (g *grid) getAllPoints() (points []point) {
	width, height := g.getDimensions()
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			points = append(points, point{x: col, y: row})
		}
	}
	return
}

func (g *grid) getValue(p point) int {
	return g[p.y][p.x]
}

func (g *grid) setValue(p point, value int) {
	g[p.y][p.x] = value
}

func (g *grid) step() (flashes int) {
	queue := queue{}
	points := g.getAllPoints()
	queue = append(queue, points...)

	for len(queue) > 0 {
		p := queue.dequeue()
		curr := g.getValue(p)
		new := curr + 1
		g.setValue(p, new)
		if new == 10 {
			flashes++
			for _, n := range getNeighbors(*g, p) {
				queue.enqueue(n)
			}
		}
	}

	for _, p := range g.getAllPoints() {
		curr := g.getValue(p)
		if curr >= 10 {
			g.setValue(p, 0)
		}
	}
	return
}

func getInput(path string) (grid grid) {
	lineStrings := utils.GetLines(path, "\n")
	for rowIndex, lineStr := range lineStrings {
		chars := strings.Split(lineStr, "")
		for colIndex, char := range chars {
			grid[rowIndex][colIndex] = utils.ParseInt(char)
		}
	}
	return
}

func getNeighbors(grid grid, curr point) (points []point) {
	width, height := grid.getDimensions()
	for _, dir := range directions {
		next := point{x: curr.x + dir.x, y: curr.y + dir.y}
		if next.x >= 0 && next.x < width && next.y >= 0 && next.y < height {
			points = append(points, next)
		}
	}
	return
}

type queue []point

func (q *queue) enqueue(p point) {
	*q = append(*q, p)
}

func (q *queue) dequeue() (p point) {
	p = (*q)[0]
	*q = (*q)[1:]
	return
}

func part1(path string) (flashes int) {
	grid := getInput(path)
	totalSteps := 100
	for step := 0; step < totalSteps; step++ {
		flashes += grid.step()
	}
	return
}

func part2(path string) (step int) {
	grid := getInput(path)
	width, height := grid.getDimensions()
	step = 0
	for {
		step++
		flashes := grid.step()
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
