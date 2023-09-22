package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

type heightmap [][]int

func (h heightmap) getDimensions() (width, height int) {
	width = len(h[0])
	height = len(h)
	return
}

func getInput(path string) (heightmap heightmap) {
	lineStrings := utils.GetLines(path, "\n")
	for _, lineStr := range lineStrings {
		line := []int{}
		chars := strings.Split(lineStr, "")
		for _, char := range chars {
			line = append(line, utils.ParseInt(char))
		}
		heightmap = append(heightmap, line)
	}
	return
}

func sum(slice []int) (total int) {
	for _, el := range slice {
		total += el
	}
	return
}

func part1(path string) int {
	heightmap := getInput(path)
	width, height := heightmap.getDimensions()
	riskLevels := []int{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			current := heightmap[y][x]

			if x-1 >= 0 {
				left := heightmap[y][x-1]
				if current >= left {
					continue
				}
			}

			if x+1 <= width-1 {
				right := heightmap[y][x+1]
				if current >= right {
					continue
				}
			}

			if y-1 >= 0 {
				top := heightmap[y-1][x]
				if current >= top {
					continue
				}
			}

			if y+1 <= height-1 {
				bottom := heightmap[y+1][x]
				if current >= bottom {
					continue
				}
			}

			riskLevel := current + 1
			riskLevels = append(riskLevels, riskLevel)
		}
	}

	return sum(riskLevels)
}

type point struct {
	x, y int
}

type direction struct {
	x, y int
}

var directions = [4]direction{
	{y: -1, x: 0}, // up
	{y: 0, x: 1},  // right
	{y: 1, x: 0},  // down
	{y: 0, x: -1}, // left
}

func getNeighbors(heightmap heightmap, curr point) (points []point) {
	width, height := heightmap.getDimensions()
	for _, dir := range directions {
		next := point{x: curr.x + dir.x, y: curr.y + dir.y}
		if next.x >= 0 && next.x < width && next.y >= 0 && next.y < height {
			points = append(points, next)
		}
	}

	return
}

const MAX_BASIN_HEIGHT = 8

func part2(path string) int {
	heightmap := getInput(path)
	width, height := heightmap.getDimensions()
	largest := []int{0, 0, 0}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			startPoint := point{x: x, y: y}
			queue := utils.Queue[point]{}
			queue.Enqueue(startPoint)
			basin := map[point]bool{}

			for queue.Len > 0 {
				currPoint, _ := queue.Dequeue()
				if basin[*currPoint] {
					continue
				}
				basin[*currPoint] = true

				for _, nextPoint := range getNeighbors(heightmap, *currPoint) {
					currHeight := heightmap[currPoint.y][currPoint.x]
					nextHeight := (heightmap)[nextPoint.y][nextPoint.x]

					if currHeight < nextHeight && nextHeight <= MAX_BASIN_HEIGHT {
						queue.Enqueue(nextPoint)
					}
				}
			}

			total := len(basin)
			if total > largest[0] {
				largest[0] = total
				sort.Ints(largest)
			}
		}
	}

	return largest[0] * largest[1] * largest[2]
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
