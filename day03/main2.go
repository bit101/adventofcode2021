package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getInput2()

	most, least := split(input, 0)

	place := 1
	for len(most) > 1 {
		most, _ = split(most, place)
		place++
	}

	place = 1
	for len(least) > 1 {
		_, least = split(least, place)
		place++
	}

	oxygen, err := strconv.ParseInt(most[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	c02, err := strconv.ParseInt(least[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("oxygen = %+v\n", oxygen)
	fmt.Printf("c02 = %+v\n", c02)

	lsRating := oxygen * c02
	fmt.Printf("lsRating = %+v\n", lsRating)

}

func split(input []string, place int) (most []string, least []string) {
	ones := []string{}
	zeros := []string{}
	for _, line := range input {
		c := line[place]
		if c == '1' {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}
	if len(ones) >= len(zeros) {
		return ones, zeros
	}
	return zeros, ones
}

////////////////////
// UTILS
////////////////////

func getInput2() []string {
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
