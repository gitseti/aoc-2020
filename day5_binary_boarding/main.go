package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var seatIDs []int = make([]int, len(lines))
	for i, line := range lines {
		rowSpecifier := line[0:7]
		columnSpecifier := line[7:10]

		bits := strings.ReplaceAll(rowSpecifier, "B", "1")
		bits = strings.ReplaceAll(bits, "F", "0")
		rowConv, _ := strconv.ParseInt(bits, 2, 64)
		row := int(rowConv)

		bits = strings.ReplaceAll(columnSpecifier, "R", "1")
		bits = strings.ReplaceAll(bits, "L", "0")
		columnConv, _ := strconv.ParseInt(bits, 2, 64)
		column := int(columnConv)

		var seatID int = row*8 + column
		seatIDs[i] = seatID
	}

	sort.Ints(seatIDs)

	fmt.Println("Part I:", seatIDs[len(seatIDs)-1])

	for i := 1; i < len(seatIDs)-1; i++ {
		if seatIDs[i-1] != seatIDs[i]-1 {
			fmt.Println("Part II:", seatIDs[i]-1)
		}
	}
}
