package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	lines := getLines(input)
	grid := NewGrid()

	h := getHorizontal(lines)

	for _, line := range h {
		count := intAbs(line.x0 - line.x1)
		dx := (line.x1 - line.x0) / count
		x, y := line.x0, line.y0
		for i := 0; i <= count; i++ {
			grid.addPoint(x, y)
			x += dx
		}
	}

	v := getVertical(lines)
	for _, line := range v {
		count := intAbs(line.y0 - line.y1)
		dy := (line.y1 - line.y0) / count
		x, y := line.x0, line.y0
		for i := 0; i <= count; i++ {
			grid.addPoint(x, y)
			y += dy
		}
	}
	/////////////////////////////
	// for part 2
	/////////////////////////////
	doPart2 := true
	if doPart2 {
		d := getDiagonal(lines)
		for _, line := range d {
			count := intAbs(line.x0 - line.x1)
			dx := (line.x1 - line.x0) / count
			dy := (line.y1 - line.y0) / count
			x := line.x0
			y := line.y0
			for i := 0; i <= count; i++ {
				grid.addPoint(x, y)
				x += dx
				y += dy
			}
		}
	}
	/////////////////////////////
	// end part 2
	/////////////////////////////

	total := 0
	for _, count := range grid.grid {
		if count > 1 {
			total++
		}
	}
	fmt.Printf("count = %+v\n", total)
}

////////////////////
// GRID
////////////////////
type Grid struct {
	grid map[string]int
}

func NewGrid() *Grid {
	return &Grid{
		grid: map[string]int{},
	}
}

func (g *Grid) addPoint(x, y int) {
	key := fmt.Sprintf("%d-%d", x, y)
	val, ok := g.grid[key]
	if !ok {
		g.grid[key] = 1
	} else {
		g.grid[key] = val + 1
	}
}

func (g *Grid) getCount(x, y int) int {
	key := fmt.Sprintf("%d-%d", x, y)
	val, ok := g.grid[key]
	if ok {
		return val
	}
	return 0
}

////////////////////
// LINES
////////////////////
type Line struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

func (l *Line) String() string {
	return fmt.Sprintf("string: [%d, %d, %d, %d]\n", l.x0, l.y0, l.x1, l.y1)
}

func NewLine(line string) *Line {
	points := strings.Split(line, "->")
	point1 := strings.Split(points[0], ",")
	x0, y0 := getPoint(point1)
	point2 := strings.Split(points[1], ",")
	x1, y1 := getPoint(point2)
	return &Line{x0, y0, x1, y1}
}

func getPoint(p []string) (int, int) {
	x, err := strconv.ParseInt(strings.TrimSpace(p[0]), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseInt(strings.TrimSpace(p[1]), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(x), int(y)
}

func getLines(input []string) []*Line {
	lines := []*Line{}
	for _, l := range input {
		line := NewLine(l)
		lines = append(lines, line)
	}
	return lines
}

////////////////////
// UTIL
////////////////////

func getHorizontal(lines []*Line) []*Line {
	out := []*Line{}
	for _, line := range lines {
		if line.y0 == line.y1 {
			out = append(out, line)
		}
	}
	return out
}

func getVertical(lines []*Line) []*Line {
	out := []*Line{}
	for _, line := range lines {
		if line.x0 == line.x1 {
			out = append(out, line)
		}
	}
	return out
}

func getDiagonal(lines []*Line) []*Line {
	out := []*Line{}
	for _, line := range lines {
		if intAbs(line.x0-line.x1) == intAbs(line.y0-line.y1) {
			out = append(out, line)
		}
	}
	return out
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInput() []string {
	input := []string{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
