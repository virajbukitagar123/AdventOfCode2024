package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convertToInt(in string) int {
	digit, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	return digit
}

func calculateResult(enabled bool, mulMatch string) int64 {
	if enabled {
		match, _ := strings.CutPrefix(mulMatch, "mul(")
		newMatch, _ := strings.CutSuffix(match, ")")
		match = newMatch
		digits := strings.Split(match, ",")
		// fmt.Println(digits)
		firstDigit := convertToInt(string(digits[0]))
		secondDigit := convertToInt(string(digits[1]))
		return int64(firstDigit) * int64(secondDigit)
	}

	return 0
}

func part1() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	pattern := `(mul\(\d+,\d+\))`
	var result int64

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	for scanner.Scan() {
		s := scanner.Text()
		matches := re.FindAllString(s, -1)
		for _, match := range matches {
			result += calculateResult(true, match)
		}
	}

	fmt.Println("Result 1 is: ", result)
}

func part2() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	pattern := `(do\(\))|(mul\(\d+,\d+\))|(don't\(\))`
	var result int64

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	enabled := true

	for scanner.Scan() {
		s := scanner.Text()
		matches := re.FindAllString(s, -1)
		for _, match := range matches {
			// fmt.Println(match)
			switch {
			case strings.Contains(match, "don't"):
				// fmt.Println("Found Don't - Enabled False")
				enabled = false
			case strings.Contains(match, "do"):
				// fmt.Println("Found Do - Enabled True")
				enabled = true
			case strings.Contains(match, "mul"):
				result += calculateResult(enabled, match)
			}
		}
	}

	fmt.Println("Result 2 is: ", result)
}

func main() {
	part1()
	part2()
	DayThree()
}
