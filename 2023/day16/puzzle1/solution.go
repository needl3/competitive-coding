package main

import (
	"errors"
	"fmt"
	"os"
)

// Read all lines
// Push to queue
// While queue
// Pop
// Is | or - or / or \, add to queue possible next coords
// Move through until nil pointer
// Add it to a map
// Increase a counter

func main() {
	// file, _ := os.Open("../input.txt")
	file, _ := os.Open("../input-submit.txt")

	inp := []string{}
	visualize := [][]VisualizedPoint{}
	for line, err := Readline(file); err == nil; line, err = Readline(file) {
		inp = append(inp, line)
		visualize = append(visualize, make([]VisualizedPoint, len(line)))
	}

	q := []Point{}
	q = append(q, Point{
		x:  0,
		y:  0,
		to: 'r',
	})

	for lptr := 0; ; lptr++ {
		if lptr >= len(q) {
			break
		}
		newQ := traverse(inp, visualize, q[lptr])
		q = append(q, newQ...)

	}

	count := printMat(visualize, true)
	fmt.Println("Sum: ", count)
}

type Point struct {
	x  int
	y  int
	to rune // r, l, u, d => Represents direction of travel
}

func (p Point) isValid(inp []string) bool {
	return p.x >= 0 && p.x < len(inp) && p.y >= 0 && p.y < len(inp[p.x])
}

type VisualizedPoint struct {
	u bool
	d bool
	l bool
	r bool
}

func (v VisualizedPoint) update(to rune) (VisualizedPoint, error) {
	switch to {
	case 'u':
		if !v.u {
			v.u = true
		} else {
			return v, errors.New("ALready set")
		}
		break
	case 'd':
		if !v.d {
			v.d = true
		} else {
			return v, errors.New("ALready set")
		}
		break
	case 'r':
		if !v.r {
			v.r = true
		} else {
			return v, errors.New("ALready set")
		}
		break
	case 'l':
		if !v.l {
			v.l = true
		} else {
			return v, errors.New("ALready set")
		}
		break
	}
	return v, nil
}

func traverse(inp []string, vis [][]VisualizedPoint, p Point) []Point {
	fmt.Println("Q: ", p.x, p.y, string(p.to))
	newQ := []Point{}

	for len(newQ) == 0 && p.isValid(inp) {
		goThrough := false
		switch inp[p.x][p.y] {
		case '-':
			if p.to != 'l' && p.to != 'r' {
				if p.y < len(inp[p.x])-1 && !vis[p.x][p.y+1].r {
					newQ = append(newQ, Point{
						x:  p.x,
						y:  p.y + 1,
						to: 'r',
					})
				}
				if p.y > 0 && !vis[p.x][p.y-1].l {
					newQ = append(newQ, Point{
						x:  p.x,
						y:  p.y - 1,
						to: 'l',
					})
				}
			} else {
				goThrough = true
			}
			break
		case '|':
			if p.to != 'u' && p.to != 'd' {
				if p.x < len(inp)-1 && !vis[p.x+1][p.y].d {
					newQ = append(newQ, Point{
						x:  p.x + 1,
						y:  p.y,
						to: 'd',
					})
				}
				if p.x > 0 && !vis[p.x-1][p.y].u {
					newQ = append(newQ, Point{
						x:  p.x - 1,
						y:  p.y,
						to: 'u',
					})
				}
			} else {
				goThrough = true
			}
			break
		case '\\':
			if p.to == 'u' && p.y > 0 && !vis[p.x][p.y-1].l {
				newQ = append(newQ, Point{
					x:  p.x,
					y:  p.y - 1,
					to: 'l',
				})
			} else if p.to == 'r' && p.x < len(inp)-1 && !vis[p.x+1][p.y].d {
				newQ = append(newQ, Point{
					x:  p.x + 1,
					y:  p.y,
					to: 'd',
				})
			} else if p.to == 'l' && p.x > 0 && !vis[p.x-1][p.y].u {
				newQ = append(newQ, Point{
					x:  p.x - 1,
					y:  p.y,
					to: 'u',
				})
			} else if p.to == 'd' && p.y < len(inp[p.x])-1 && !vis[p.x][p.y+1].r {
				newQ = append(newQ, Point{
					x:  p.x,
					y:  p.y + 1,
					to: 'r',
				})
			}
			break
		case '/':
			if p.to == 'd' && p.y > 0 && !vis[p.x][p.y-1].l {
				newQ = append(newQ, Point{
					x:  p.x,
					y:  p.y - 1,
					to: 'l',
				})
			} else if p.to == 'l' && p.x < len(inp)-1 && !vis[p.x+1][p.y].d {
				newQ = append(newQ, Point{
					x:  p.x + 1,
					y:  p.y,
					to: 'd',
				})
			} else if p.to == 'r' && p.x > 0 && !vis[p.x-1][p.y].u {
				newQ = append(newQ, Point{
					x:  p.x - 1,
					y:  p.y,
					to: 'u',
				})
			} else if p.to == 'u' && p.y < len(inp[p.x])-1 && !vis[p.x][p.y+1].r {
				newQ = append(newQ, Point{
					x:  p.x,
					y:  p.y + 1,
					to: 'r',
				})
			}
			break
		default:
			goThrough = true
		}
		if nView, err := vis[p.x][p.y].update(p.to); err == nil {
			vis[p.x][p.y] = nView
		} else {
			break
		}
		if goThrough {
			switch p.to {
			case 'l':
				p.y--
				break
			case 'r':
				p.y++
				break
			case 'u':
				p.x--
				break
			case 'd':
				p.x++
				break
			}
		}
	}

	return newQ
}

func printMat(arr2d [][]VisualizedPoint, display bool) int {
	count := 0
	for i := range arr2d {
		for j := range arr2d[i] {
			if arr2d[i][j].d || arr2d[i][j].u || arr2d[i][j].l || arr2d[i][j].r {
				count++
				if display {
					fmt.Printf("#")
				}
			} else {
				if display {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Println()
	}
	return count
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
