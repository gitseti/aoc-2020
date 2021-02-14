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

func parseInput() ([]int, []int, []byte, []string) {
	data, err := ioutil.ReadFile("input.txt")
	check(err)

	var lines []string = strings.Split(string(data), "\n")
	var minOccurences = []int{}
	var maxOccurences = []int{}
	var charsToCheck = []byte{}
	var passwords = []string{}

	for _, line := range lines {

		minOcc, err := strconv.Atoi(line[0:strings.IndexRune(line, '-')])
		check(err)

		maxOcc, err := strconv.Atoi(line[strings.IndexRune(line, '-')+1 : strings.IndexRune(line, ' ')])
		check(err)

		charToCheck := line[strings.IndexRune(line, ' ')+1]
		check(err)

		password := strings.Split(line, ": ")[1]

		minOccurences = append(minOccurences, minOcc)
		maxOccurences = append(maxOccurences, maxOcc)
		charsToCheck = append(charsToCheck, charToCheck)
		passwords = append(passwords, password)
	}
	return minOccurences, maxOccurences, charsToCheck, passwords
}

func main() {
	minOccurences, maxOccurences, charsToCheck, passwords := parseInput()

	size := len(passwords)
	var validPasswordsPart1 int = 0
	var validPasswordsPart2 int = 0

	for i := 0; i < size; i++ {
		minOcc := minOccurences[i]
		maxOcc := maxOccurences[i]
		char := charsToCheck[i]
		password := passwords[i]

		var occs int = 0

		for j := 0; j < len(password); j++ {
			if password[j] == char {
				occs++
			}
		}

		if occs >= minOcc && occs <= maxOcc {
			validPasswordsPart1++
		}

		if (password[minOcc-1] != char && password[maxOcc-1] == char) || (password[minOcc-1] == char && password[maxOcc-1] != char) {
			validPasswordsPart2++
		}

	}

	fmt.Printf("Valid Password for Part I: %d\n", validPasswordsPart1)
	fmt.Printf("Valid Password for Part II: %d\n", validPasswordsPart2)
}
