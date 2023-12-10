package main

import (
	"errors"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func (p Point) coord() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")
	var input []string
	for {
		if line, err := Readline(file); err == nil {
			input = append(input, line)
		} else {
			break
		}
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'S' {
				currPoint := &Point{
					X: i,
					Y: j,
				}

				var startPoint Point
				if j+1 < len(input[i]) && input[i][j+1] == '-' || input[i][j+1] == 'J' || input[i][j+1] == '7' {
					startPoint.X = i
					startPoint.Y = j + 1
				} else if j-1 >= 0 && input[i][j-1] == '-' || input[i][j-1] == 'F' || input[i][j-1] == 'L' {
					startPoint.X = i
					startPoint.Y = j - 1
				} else if i-1 >= 0 && input[i-1][j] == '-' || input[i-1][j] == 'F' || input[i-1][j] == 'L' {
					startPoint.X = i - 1
					startPoint.Y = j
				} else {
					startPoint.X = i + 1
					startPoint.Y = j
				}
				if n, err := findPath(&startPoint, currPoint, input); err == nil { // Up
					fmt.Println("Final path length: ", n)
					return
				}
			}
		}
	}
}

func findPath(currentPoint *Point, origin *Point, input []string) (int, error) {
	if currentPoint.X < 0 || currentPoint.X >= len(input[0]) || currentPoint.Y < 0 || currentPoint.Y >= len(input) {
		return 0, errors.New("OFB")
	}
	length := 2 // S and this starting cell counts
	for {
		if nxt, err := nextPoint(currentPoint, origin, input); err != nil {
			return 0, errors.New("No loop")
		} else if input[nxt.X][nxt.Y] == 'S' {
			break
		} else {
			origin = currentPoint
			currentPoint = nxt
			length++
		}
	}
	fmt.Println()

	return length / 2, nil
}

func nextPoint(currPoint *Point, origin *Point, input []string) (*Point, error) {
	// fmt.Printf("%s", string(input[currPoint.X][currPoint.Y]))
	var points []Point
	switch input[currPoint.X][currPoint.Y] {
	case '|':
		if currPoint.X-1 >= 0 {
			points = append(points, Point{
				X: currPoint.X - 1,
				Y: currPoint.Y,
			})
		}
		if currPoint.X+1 < len(input) {
			points = append(points, Point{
				X: currPoint.X + 1,
				Y: currPoint.Y,
			})
		}
		break
	case '-':
		if currPoint.Y-1 >= 0 {
			points = append(points, Point{
				X: currPoint.X,
				Y: currPoint.Y - 1,
			})
		}
		if currPoint.Y+1 < len(input[currPoint.X]) {
			points = append(points, Point{
				X: currPoint.X,
				Y: currPoint.Y + 1,
			})
		}
		break
	case 'L':
		if currPoint.X-1 >= 0 {
			points = append(points, Point{
				X: currPoint.X - 1,
				Y: currPoint.Y,
			})
		}
		if currPoint.Y+1 < len(input[currPoint.X]) {
			points = append(points, Point{
				X: currPoint.X,
				Y: currPoint.Y + 1,
			})
		}
		break
	case 'J':
		if currPoint.X-1 >= 0 {
			points = append(points, Point{
				X: currPoint.X - 1,
				Y: currPoint.Y,
			})
		}
		if currPoint.Y-1 >= 0 {
			points = append(points, Point{
				X: currPoint.X,
				Y: currPoint.Y - 1,
			})
		}
		break
	case '7':
		if currPoint.X+1 < len(input) {
			points = append(points, Point{
				Y: currPoint.Y,
				X: currPoint.X + 1,
			})
		}
		if currPoint.Y-1 >= 0 {
			points = append(points, Point{
				Y: currPoint.Y - 1,
				X: currPoint.X,
			})
		}
		break
	case 'F':
		if currPoint.X+1 < len(input) {
			points = append(points, Point{
				Y: currPoint.Y,
				X: currPoint.X + 1,
			})
		}
		if currPoint.Y+1 < len(input[currPoint.X]) {
			points = append(points, Point{
				Y: currPoint.Y + 1,
				X: currPoint.X,
			})
		}
		break
	default:
		return nil, errors.New("No pipe")
	}

	for _, point := range points {
		if point.coord() != origin.coord() {
			return &point, nil
		}
	}
	return nil, errors.New("No pipe")
}

func Readline(f *os.File) (string, error) {
	var line []byte
	singleFileContent := make([]byte, 1)
	i := 0
	for {
		if read, err := f.Read(singleFileContent); err == nil && read > 0 {
			if singleFileContent[0] != byte('\n') {
				line = append(line, singleFileContent...)
			} else {
				break
			}
		} else {
			return "", errors.New("eof")
		}
		i = i + 1
	}
	return string(line), nil
}
