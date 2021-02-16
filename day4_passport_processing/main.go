package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var validHcl = regexp.MustCompile("^#[0-9a-f]{6}$")
var validPid = regexp.MustCompile("^[0-9]{9}$")
var validEyeColor = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")

func parseInput() []map[string]string {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	passportData := make([]map[string]string, 0)

	ppIndex := 0
	passportData = append(passportData, make(map[string]string))
	for _, line := range lines {
		if line != "" {
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				keyValuePair := strings.Split(pair, ":")
				if len(keyValuePair) == 2 {
					passportData[ppIndex][keyValuePair[0]] = keyValuePair[1]
				}
			}
		} else {
			ppIndex++
			passportData = append(passportData, make(map[string]string))
		}
	}

	if len(passportData[ppIndex]) == 0 {
		passportData = passportData[0 : ppIndex-1]
	}

	return passportData
}

func checkInRange(toCheck string, min int, max int) bool {
	value, err := strconv.Atoi(toCheck)
	if err != nil {
		panic(err)
	}
	if !(value >= min && value <= max) {
		return false
	}
	return true
}

func filterPartOne(passportData []map[string]string) []map[string]string {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	filteredMap := make([]map[string]string, 0)

	for _, passport := range passportData {
		valid := true
		for _, requiredField := range requiredFields {
			if _, ok := passport[requiredField]; !ok {
				valid = false
				break
			}
		}
		if valid {
			filteredMap = append(filteredMap, passport)
		}
	}

	return filteredMap
}

func filterPartTwo(passportData []map[string]string) []map[string]string {

	filteredMap := make([]map[string]string, 0)

	for _, passport := range passportData {
		if checkValid(passport) {
			filteredMap = append(filteredMap, passport)
		}
	}

	return filteredMap
}

func checkValid(passport map[string]string) bool {

	if !checkInRange(passport["byr"], 1920, 2002) || len(passport["byr"]) != 4 {
		return false
	}

	if !checkInRange(passport["iyr"], 2010, 2020) || len(passport["iyr"]) != 4 {
		return false
	}

	if !checkInRange(passport["eyr"], 2020, 2030) || len(passport["eyr"]) != 4 {
		return false
	}

	if passport["hgt"] == "118" {
		fmt.Println()
	}
	hgt := passport["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		if !checkInRange(strings.Split(hgt, "c")[0], 150, 193) {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		if !checkInRange(strings.Split(hgt, "i")[0], 59, 76) {
			return false
		}
	} else {
		return false
	}

	if !validEyeColor.MatchString(passport["ecl"]) {
		return false
	}

	if !validHcl.MatchString(passport["hcl"]) {
		return false
	}

	if !validPid.MatchString(passport["pid"]) {
		return false
	}

	return true
}

func main() {

	passportData := parseInput()

	passportData = filterPartOne(passportData)

	fmt.Printf("Part I : Found %d valid passports\n", len(passportData))

	passportData = filterPartTwo(passportData)

	fmt.Printf("Part II : Found %d valid passports\n", len(passportData))

}
