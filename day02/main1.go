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
	sub := NewSub()

	input := getInput()
	for _, cmd := range input {
		sub.Execute(cmd)
	}
	fmt.Printf("depth = %+v\n", sub.Depth)
	fmt.Printf("pos = %+v\n", sub.Position)
	product := sub.Depth * sub.Position
	fmt.Printf("product = %+v\n", product)
}

////////////////////
// COMMAND
////////////////////
type Command struct {
	Dir    string
	Amount int
}

func NewCommand(line string) *Command {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		log.Fatalf("'%s' is not a valid command", line)
	}
	dir := parts[0]
	amount, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return &Command{
		dir,
		int(amount),
	}
}

////////////////////
// SUB
////////////////////
type Sub struct {
	Depth    int
	Position int
}

func NewSub() *Sub {
	return &Sub{0, 0}
}

func (s *Sub) Execute(cmd *Command) {
	switch cmd.Dir {
	case "forward":
		s.Position += cmd.Amount
	case "down":
		s.Depth += cmd.Amount
	case "up":
		s.Depth -= cmd.Amount
	}
}

////////////////////
// UTIL
////////////////////
func getInput() []*Command {
	input := []*Command{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := NewCommand(line)
		input = append(input, cmd)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
