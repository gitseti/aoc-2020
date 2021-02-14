package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcTreeHits(lines []string, right int, down int) int {
	var patternSize = len(lines[0])
	var hitTrees int = 0
	var j int = 0
	for i := 0; i < len(lines); i += down {
		if lines[i][j] == '#' {
			hitTrees++
		}

		if j+right >= patternSize {
			j = j + right - patternSize
		} else {
			j += right
		}
	}

	return hitTrees
}

func main() {

	data, err := ioutil.ReadFile("input.txt")
	check(err)

	var lines []string = strings.Split(string(data), "\n")

	/*

		right 3, down 1

	*/
	partOneSolution := calcTreeHits(lines, 3, 1)
	fmt.Printf("Part I: %d\n", partOneSolution)

	/*

	   Right 1, down 1.
	   Right 3, down 1.
	   Right 5, down 1.
	   Right 7, down 1.
	   Right 1, down 2.

	   What do you get if you multiply together the number of trees encountered on each of the listed slopes?

	*/
	partTwoSolution := calcTreeHits(lines, 1, 1) * calcTreeHits(lines, 3, 1) * calcTreeHits(lines, 5, 1) * calcTreeHits(lines, 7, 1) * calcTreeHits(lines, 1, 2)
	fmt.Printf("Part II: %d\n", partTwoSolution)
}
