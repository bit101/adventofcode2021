package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getInput()

	count := 0
	// part 1: window = 1
	// part 2: window = 3
	window := 3

	prevSum := windowSum(input, 0, window)
	for i := 1; i <= len(input)-window; i++ {
		sum := windowSum(input, i, window)
		if sum > prevSum {
			count++
		}
		prevSum = sum
	}
	fmt.Println(count)
}

////////////////////
// UTILS
////////////////////
func windowSum(arr []int, start, length int) int {
	sum := 0
	for i := start; i < start+length; i++ {
		sum += arr[i]
	}
	return sum
}

func getInput() []int {
	input := []int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, int(num))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
