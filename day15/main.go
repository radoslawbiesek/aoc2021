package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

func getGrid(path string) (grid utils.Grid) {
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

func gridPointToGraphIndex(grid utils.Grid, point utils.Point) int {
	width, _ := grid.GetDimensions()

	return point.Y*width + point.X
}

func gridToGraph(grid utils.Grid) (graph utils.WeightedAdjList) {
	width, height := grid.GetDimensions()
	graph = make(utils.WeightedAdjList, width*height)

	for _, currPoint := range grid.GetAllPoints() {
		currPointIdx := gridPointToGraphIndex(grid, currPoint)

		for _, neighbor := range utils.Get4Neighbors(grid, currPoint) {
			graphEdge := utils.GraphEdge{
				From: currPointIdx,
				To:   gridPointToGraphIndex(grid, neighbor),
				Cost: grid.GetValue(neighbor),
			}

			if !utils.Has(graph[currPointIdx], graphEdge) {
				graph[currPointIdx] = append(graph[currPointIdx], graphEdge)
			}

			graphReturnEdge := utils.GraphEdge{
				From: graphEdge.To,
				To:   currPointIdx,
				Cost: grid.GetValue(currPoint),
			}

			if !utils.Has(graph[graphEdge.To], graphReturnEdge) {
				graph[graphEdge.To] = append(graph[graphEdge.To], graphReturnEdge)
			}
		}
	}

	return
}

func part1(path string) int {
	grid := getGrid(path)
	graph := gridToGraph(grid)
	risk, _ := utils.DijkstraList(graph, 0, len(graph)-1)

	return risk
}

func main() {
	fmt.Println("Test input: ")
	fmt.Printf("Part 1: %d\n", part1("./test-input.txt"))
	fmt.Println("")
	fmt.Println("Input: ")
	fmt.Printf("Part 1: %d\n", part1("./input.txt"))
}
