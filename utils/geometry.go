package utils

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type Line struct {
	P1 Point
	P2 Point
}

func (l Line) IsVertical() bool {
	return l.P1.X == l.P2.X
}

func (l Line) IsHorizontal() bool {
	return l.P1.Y == l.P2.Y
}

func (line Line) GetPoints() (points []Point) {
	diffX := line.P2.X - line.P1.X
	diffY := line.P2.Y - line.P1.Y
	stepsCount := int(math.Max(math.Abs(float64(diffX)), math.Abs(float64(diffY))))

	var stepX, stepY int
	if diffX == 0 {
		stepX = 0
	} else if diffX > 0 {
		stepX = 1
	} else {
		stepX = -1
	}

	if diffY == 0 {
		stepY = 0
	} else if diffY > 0 {
		stepY = 1
	} else {
		stepY = -1
	}

	for i := 0; i <= stepsCount; i++ {
		x := line.P1.X + i*stepX
		y := line.P1.Y + i*stepY
		point := Point{X: x, Y: y}
		points = append(points, point)
	}

	return
}

type Diagram map[string]int

func getDiagramKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func (d Diagram) GetPositionValue(x, y int) int {
	key := getDiagramKey(x, y)
	return d[key]
}

func (d *Diagram) SetPositionValue(x, y, value int) {
	key := getDiagramKey(x, y)
	(*d)[key] = value
}
