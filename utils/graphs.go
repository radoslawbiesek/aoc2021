package utils

import (
	"fmt"
	"math"
	"sort"
)

type GraphEdge struct {
	From, To, Cost int
}

func (g GraphEdge) String() string {
	return fmt.Sprintf("{ From: %d, to: %d, cost: %d }", g.From, g.To, g.Cost)
}

type WeightedAdjList [][]GraphEdge

func hasUnvisited(seen []bool, dists []float64) bool {
	for i := range seen {
		if seen[i] {
			continue
		}

		if dists[i] < math.Inf(1) {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	lowestDist := math.Inf(1)

	for i := range seen {
		if seen[i] {
			continue
		}

		if lowestDist > dists[i] {
			idx = i
			lowestDist = dists[i]
		}
	}

	return idx
}

func DijkstraList(graph WeightedAdjList, source, target int) (totalCost int, path []int) {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	dists := make([]float64, len(graph))
	for i := range dists {
		dists[i] = math.Inf(1)
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		if curr < 0 {
			break
		}

		seen[curr] = true
		adjs := graph[curr]

		for _, edge := range adjs {
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + float64(edge.Cost)
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	totalCost = int(dists[target])

	curr := target
	path = append(path, target)
	for prev[curr] != -1 {
		path = append(path, prev[curr])
		curr = prev[curr]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(path)))

	return
}
