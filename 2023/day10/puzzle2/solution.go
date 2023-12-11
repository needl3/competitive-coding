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

func printMat(mat [][]string) {
	for _, i := range mat {
		for _, j := range i {
			fmt.Printf("%v", j)
		}
		fmt.Println()
	}
}

func main() {
	// file, _ := os.Open("input.txt")
	file, _ := os.Open("../input-submit.txt")

	var input []string
	for {
		if line, err := Readline(file); err == nil {
			input = append(input, line)
		} else {
			break
		}
	}

	startPoint, origin := findStartPoint(input)
	if n, loop, finalPoint, err := findPath(startPoint, origin, input); err == nil {
		// printMat(loop)
		loop[origin.X][origin.Y] = findOriginChar(origin, startPoint, finalPoint)
		fmt.Printf(loop[origin.X][origin.Y])
		count := 0
		for x := 0; x < len(loop); x++ {
			state := 0
			for y := 0; y < len(loop[0]); y++ {
				state, count = updateCount(state, count, loop[x][y])
				// fmt.Printf("%d", state)
			}
			// fmt.Println()
		}
		fmt.Println("Area: ", count, n)
		return
	}
}

func updateCount(state int, count int, input string) (int, int) {
	switch fmt.Sprintf("%d%s", state, input) {
	case "0|":
		return 1, count
	case "0F":
		return 2, count
	case "0L":
		return 3, count
	case "1|":
		return 0, count
	case "1F":
		return 4, count
	case "1L":
		return 5, count
	case "1 ":
		return 1, count + 1
	case "2J":
		return 1, count
	case "27":
		return 0, count
	case "37":
		return 1, count
	case "3J":
		return 0, count
	case "47":
		return 1, count
	case "4J":
		return 0, count
	case "5J":
		return 1, count
	case "57":
		return 0, count
	default:
		return state, count
	}
}

func findOriginChar(origin *Point, sPoint *Point, fPoint *Point) string {
	fmt.Println(sPoint, fPoint, origin)
	if (sPoint.Y-fPoint.Y)%2 == 0 && sPoint.Y-fPoint.Y != 0 {
		return "-"
	} else if (sPoint.X-fPoint.X)%2 == 0 && sPoint.X-fPoint.X != 0 {
		return "|"
	} else {
		rl := 0
		ud := 0
		if origin.Y-sPoint.Y > 0 {
			rl = -1 // left
		} else {
			rl = 1 // right
		}
		if origin.X-fPoint.X > 0 {
			ud = 1 // up
		} else {
			ud = -1 // down
		}

		switch fmt.Sprintf("%d%d", rl, ud) {
		case "-1-1":
			return "7"
		case "-11":
			return "J"
		case "11":
			return "L"
		default:
			return "F"
		}
	}
}

func findStartPoint(input []string) (*Point, *Point) {
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
				return &startPoint, currPoint
			}
		}
	}
	return nil, nil
}

func findPath(currentPoint *Point, origin *Point, input []string) (int, [][]string, *Point, error) {
	if currentPoint.X < 0 || currentPoint.X >= len(input) || currentPoint.Y < 0 || currentPoint.Y >= len(input[0]) {
		return 0, nil, nil, errors.New("OFB")
	}
	length := 2 // S and this starting cell counts

	// Initialize a false matrix
	loop := make([][]string, len(input))
	for i := range loop {
		loop[i] = make([]string, len(input[0]))
		for j := range loop[i] {
			loop[i][j] = " "
		}
	}

	loop[origin.X][origin.Y] = "S"
	loop[currentPoint.X][currentPoint.Y] = string(input[currentPoint.X][currentPoint.Y])

	var finalPoint *Point
	for {
		if nxt, err := nextPoint(currentPoint, origin, input); err != nil {
			return 0, nil, nil, errors.New("No loop")
		} else if input[nxt.X][nxt.Y] == 'S' {
			finalPoint = currentPoint
			break
		} else {
			loop[nxt.X][nxt.Y] = string(input[nxt.X][nxt.Y])
			origin = currentPoint
			currentPoint = nxt
			length++
		}
	}
	fmt.Println()

	return length / 2, loop, finalPoint, nil
}

func nextPoint(currPoint *Point, origin *Point, input []string) (*Point, error) {
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
