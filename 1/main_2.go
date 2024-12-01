package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var leftList []int64
	rightList := make(map[int64]int)

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

		if _, ok := rightList[int64(secondDigit)]; !ok {
			rightList[int64(secondDigit)] = 0
		}

		rightList[int64(secondDigit)] = rightList[int64(secondDigit)] + 1
	}

	var result int64 = 0

	for _, v := range leftList {
		numOcc, ok := rightList[v]
		if ok {
			result += v * int64(numOcc)
		}
	}

	fmt.Println("Result is: ", result)
}
