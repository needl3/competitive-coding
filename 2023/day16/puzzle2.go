package day16

import (
	"fmt"
	"os"

	"github.com/needl3/adventofcode/2023/utils"
)

func Puzzle2() {
	file, _ := os.Open("day16/input.txt")

	inp := []string{}
	for line, err := utils.Readline(file); err == nil; line, err = utils.Readline(file) {
		inp = append(inp, line)
	}

	maxTilesEnergized := 0
	x := 0
	y := 0
	for {
		currentPoints := gimmePoints(x, y, inp)
		for lptr := 0; lptr < len(currentPoints); lptr++ {
			visualize := [][]VisualizedPoint{}
			for i := range inp {
				visualize = append(visualize, make([]VisualizedPoint, len(inp[i])))
			}

			currentPoint := currentPoints[lptr]

			q := []Point{currentPoint}
			for lptr := 0; ; lptr++ {
				if lptr >= len(q) {
					break
				}
				newQ := traverse(inp, visualize, q[lptr])
				q = append(q, newQ...)

			}

			count := printMat(visualize, false)
			if maxTilesEnergized < count {
				maxTilesEnergized = count
			}

		}
		if x == 0 {
			if y < len(inp[x])-1 {
				y++
			} else {
				x++
				y = 0
			}
		} else if x < len(inp)-1 {
			if y == 0 {
				y = len(inp[x]) - 1
			} else {
				y = 0
				x++
			}
		} else {
			if y < len(inp[x])-1 {
				y++
			} else {
				break
			}
		}
	}
	fmt.Println("Max tiles:", maxTilesEnergized)
}

func gimmePoints(x int, y int, inp []string) []Point {
	possiblePoints := []Point{}

	if x == 0 {
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'd',
		})
	} else if x == len(inp)-1 {
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'u',
		})
	} else if x > 0 && x < len(inp)-1 {
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'u',
		})
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'd',
		})
	}

	if y == 0 {
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'r',
		})
	} else if y == len(inp[x])-1 {
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'l',
		})
	} else if (x == 0 || x == len(inp)-1) && y > 0 && y < len(inp[x])-1 {
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'l',
		})
		possiblePoints = append(possiblePoints, Point{
			x:  x,
			y:  y,
			to: 'r',
		})
	}
	return possiblePoints
}
