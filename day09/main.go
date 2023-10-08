package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

func getInput(path string) (heightmap utils.Grid) {
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
	width, height := heightmap.GetDimensions()
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

const MAX_BASIN_HEIGHT = 8

func part2(path string) int {
	heightmap := getInput(path)
	width, height := heightmap.GetDimensions()
	largest := []int{0, 0, 0}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			startPoint := utils.Point{X: x, Y: y}
			queue := utils.Queue[utils.Point]{}
			queue.Enqueue(startPoint)
			basin := map[utils.Point]bool{}

			for queue.Len > 0 {
				currPoint, _ := queue.Dequeue()
				if basin[*currPoint] {
					continue
				}
				basin[*currPoint] = true

				for _, nextPoint := range utils.GetNeighbors4(heightmap, *currPoint) {
					currHeight := heightmap[currPoint.Y][currPoint.X]
					nextHeight := (heightmap)[nextPoint.Y][nextPoint.X]

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
