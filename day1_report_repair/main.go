package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() []int {
	data, err := ioutil.ReadFile("input.txt")
	check(err)

	var lines []string = strings.Split(string(data), "\n")
	var numbers = []int{}

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		check(err)

		numbers = append(numbers, number)
	}
	return numbers
}

func main() {

	numbers := parseInput()
	size := len(numbers)

	// Naively just try out every possible combination -> O(n^3) complexity
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Printf("Found numbers: %d + %d + %d = 2020\n", numbers[i], numbers[j], numbers[k])
					fmt.Printf("Multiplying numbers: %d x %d x %d = %d\n", numbers[i], numbers[j], numbers[k], numbers[i]*numbers[j]*numbers[k])
				}
			}
		}
	}

}
