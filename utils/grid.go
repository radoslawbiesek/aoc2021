package utils

type direction struct {
	x, y int
}

var directions4 = []direction{
	{y: -1, x: 0}, // up
	{y: 0, x: 1},  // right
	{y: 1, x: 0},  // down
	{y: 0, x: -1}, // left
}

var directions8 = []direction{
	{y: -1, x: -1}, // NW
	{y: -1, x: 0},  // N
	{y: -1, x: 1},  // NE
	{y: 0, x: 1},   // E
	{y: 0, x: -1},  // W
	{y: 1, x: 0},   // S
	{y: 1, x: -1},  // SW
	{y: 1, x: 1},   // SE
}

type Grid [][]int

func (g Grid) GetDimensions() (width, height int) {
	width = len(g[0])
	height = len(g)
	return
}

func (g Grid) GetAllPoints() (points []Point) {
	width, height := g.GetDimensions()
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			points = append(points, Point{X: col, Y: row})
		}
	}
	return
}

func (g Grid) GetValue(p Point) int {
	return g[p.Y][p.X]
}

func (g *Grid) SetValue(p Point, value int) {
	(*g)[p.Y][p.X] = value
}

func getNeighbors(grid Grid, curr Point, directions []direction) (points []Point) {
	width, height := grid.GetDimensions()
	for _, dir := range directions {
		next := Point{X: curr.X + dir.x, Y: curr.Y + dir.y}
		if next.X >= 0 && next.X < width && next.Y >= 0 && next.Y < height {
			points = append(points, next)
		}
	}
	return
}

func Get4Neighbors(grid Grid, curr Point) []Point {
	return getNeighbors(grid, curr, directions4)
}

func Get8Neighbors(grid Grid, curr Point) []Point {
	return getNeighbors(grid, curr, directions8)
}
