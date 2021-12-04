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
	numbers, tables := getInput()
	findFirstWinner(numbers, tables)
	clearTables(tables)
	findLastWinner(numbers, tables)
}

func findFirstWinner(numbers []int, tables []*Table) {
	for _, num := range numbers {
		for _, table := range tables {
			table.CheckValue(num)
			table.CheckWin()
			if table.won {
				total := table.UnmarkedTotal()
				score := total * num
				fmt.Printf("first winner score = %+v\n", score)
				return
			}
		}
	}

}

func clearTables(tables []*Table) {
	for _, t := range tables {
		t.UnmarkAll()
	}
}

func findLastWinner(numbers []int, tables []*Table) {
	for _, num := range numbers {
		for _, table := range tables {
			table.CheckValue(num)
			table.CheckWin()
			if table.won && len(tables) == 1 {
				total := table.UnmarkedTotal()
				score := total * num
				fmt.Printf("last winner score = %+v\n", score)
			}
		}
		tables = getNonWinners(tables)
	}
}

func getNonWinners(tables []*Table) []*Table {
	nonWinners := []*Table{}
	for _, t := range tables {
		if !t.won {
			nonWinners = append(nonWinners, t)
		}
	}
	return nonWinners
}

////////////////////
// CELL
////////////////////
type Cell struct {
	value  int
	marked bool
}

func NewCell(value int) *Cell {
	return &Cell{value, false}
}

////////////////////
// TABLE
////////////////////
type Table struct {
	cells [25]*Cell
	won   bool
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) String() string {
	s := ""
	for i, c := range t.cells {
		if i%5 == 0 {
			s += "\n"
		}
		s += fmt.Sprintf("%d ", c.value)
	}
	s += "\n"
	return s
}

func (t *Table) SetCell(cell *Cell, index int) {
	t.cells[index] = cell
}

func (t *Table) UnmarkAll() {
	for _, c := range t.cells {
		c.marked = false
	}
	t.won = false
}

func (t *Table) CheckValue(value int) {
	for _, cell := range t.cells {
		if cell.value == value {
			cell.marked = true
		}
	}
}

func (t *Table) CheckWin() {
	if t.won {
		return
	}
	for i := 0; i < 5; i++ {
		won := true
		for j := 0; j < 5; j++ {
			index := i*5 + j
			if !t.cells[index].marked {
				won = false
			}
		}
		if won {
			t.won = true
			return
		}
	}
	for i := 0; i < 5; i++ {
		won := true
		for j := 0; j < 5; j++ {
			index := j*5 + i
			if !t.cells[index].marked {
				won = false
			}
		}
		if won {
			t.won = true
			return
		}
	}
}

func (t *Table) UnmarkedTotal() int {
	total := 0
	for _, c := range t.cells {
		if !c.marked {
			total += c.value
		}
	}
	return total
}

////////////////////
// UTIL
////////////////////

func getInput() ([]int, []*Table) {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// numbers
	numbers := getNumbers(scanner)

	// tables
	tables := getTables(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return numbers, tables
}

func getNumbers(scanner *bufio.Scanner) []int {
	scanner.Scan()
	line := scanner.Text()
	numberStrs := strings.Split(line, ",")
	numbers := []int{}
	for _, n := range numberStrs {
		num, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, int(num))
	}
	return numbers
}

func getTables(scanner *bufio.Scanner) []*Table {
	tables := []*Table{}
	for scanner.Scan() {
		table := NewTable()
		index := 0
		for i := 0; i < 5; i++ {
			scanner.Scan()
			line := scanner.Text()
			row := strings.Split(line, " ")
			for _, n := range row {
				n = strings.TrimSpace(n)
				if n == "" {
					continue
				}
				num, err := strconv.ParseInt(n, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				table.SetCell(NewCell(int(num)), index)
				index++
			}
		}
		tables = append(tables, table)
	}
	return tables
}
