package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := getInput()
	counts := [12]int{}
	half := len(input) / 2

	for _, line := range input {
		for i, c := range line {
			if c == '1' {
				counts[i]++
			}
		}
	}
	gamma := 0
	epsilon := 0
	for _, i := range counts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if i >= half {
			gamma += 1
		} else {
			epsilon += 1
		}
	}
	fmt.Printf("gamma = %+v\n", gamma)
	fmt.Printf("epsilon = %+v\n", epsilon)
	power := gamma * epsilon
	fmt.Printf("power = %+v\n", power)
}

////////////////////
// UTIL
////////////////////

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
