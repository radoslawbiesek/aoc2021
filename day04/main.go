package main

import (
	"fmt"
	"strings"

	"github.com/radoslawbiesek/aoc2021/utils"
)

type boardCell struct {
	number int
	marked bool
}
type board [5][5]boardCell

func newBoard(linesStr []string) (b board) {
	for rowIndex, rowStr := range linesStr {
		for colIndex, cellStr := range utils.Filter(strings.Split(rowStr, " "), utils.IsNonEmptyStr) {
			cell := utils.ParseInt(cellStr)
			b[rowIndex][colIndex] = boardCell{number: cell, marked: false}
		}
	}

	return
}

func (b board) String() (str string) {
	for _, row := range b {
		for _, cell := range row {
			str += fmt.Sprintf("%v", cell.number)
			if cell.marked {
				str += "x"
			}
			str += " "
		}
		str += "\n"
	}

	return
}

func (b *board) markNumber(number int) {
	for rowIndex, row := range b {
		for colIndex, cell := range row {
			if cell.number == number {
				b[rowIndex][colIndex].marked = true
				return
			}
		}
	}
}

func (b board) hasBingo() bool {
	for _, row := range b {
		bingo := true
		for _, cell := range row {
			if !cell.marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}

	for colIndex := 0; colIndex < len(b[0]); colIndex++ {
		bingo := true
		for _, row := range b {
			cell := row[colIndex]
			if !cell.marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}

	return false
}

func (b board) countUnmarked() (total int) {
	for _, row := range b {
		for _, cell := range row {
			if !cell.marked {
				total += cell.number
			}
		}
	}
	return
}

func getInput(path string) (numbers []int, boards []board) {
	lines := utils.GetLines(path, "\n")
	numbers = utils.Map(strings.Split(lines[0], ","), utils.ParseInt)
	lines = lines[2:]

	for i := 0; i < len(lines); i += 6 {
		strLines := lines[i : i+5]
		boards = append(boards, newBoard(strLines))
	}

	return
}

func part1(path string) int {
	numbers, boards := getInput(path)

	winners := []int{}
	var lastNumber int
	var bestScore int

	for _, number := range numbers {
		lastNumber = number
		for boardIndex := range boards {
			boards[boardIndex].markNumber(number)
			if boards[boardIndex].hasBingo() {
				winners = append(winners, boardIndex)
			}
		}

		if len(winners) > 0 {
			break
		}
	}

	for _, winnerIndex := range winners {
		winner := boards[winnerIndex]
		score := winner.countUnmarked()
		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore * lastNumber
}

func part2(path string) int {
	numbers, boards := getInput(path)

	winners := map[int]bool{}
	var lastNumber int
	var lastWinner board

	for _, number := range numbers {
		lastNumber = number
		for boardIndex := range boards {
			if !winners[boardIndex] {
				boards[boardIndex].markNumber(number)
				if boards[boardIndex].hasBingo() {
					winners[boardIndex] = true
					lastWinner = boards[boardIndex]
				}
			}
		}

		if len(winners) == len(boards) {
			break
		}
	}

	return lastNumber * lastWinner.countUnmarked()
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
