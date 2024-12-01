package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rightList []int64
	var leftList []int64

	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Split(s, "   ")
		firstDigit, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}

		secondDigit, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		leftList = append(leftList, int64(firstDigit))
		rightList = append(rightList, int64(secondDigit))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var result int64 = 0

	for i, _ := range leftList {
		result += abs(leftList[i] - rightList[i])
	}

	fmt.Println("Result is: ", result)
}

func abs(x int64) int64 {
	if x < 0 {
		return x * -1
	}

	return x
}
