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

func part1() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	arr := make([][]int, 0, 10)
	index := 0
	result := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Split(s, " ")
		// fmt.Println(line, len(line))
		temp := make([]int, len(line))
		arr = append(arr, temp)
		for i, v := range line {
			convValue, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			arr[index][i] = convValue
		}
		index++
	}

	//fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		if isSliceSafe(arr[i]) {
			result++
		}
	}

	fmt.Println("Result is: ", result)
}

func part2() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	arr := make([][]int, 0, 10)
	index := 0
	result := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Split(s, " ")
		// fmt.Println(line, len(line))
		temp := make([]int, len(line))
		arr = append(arr, temp)
		for i, v := range line {
			convValue, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			arr[index][i] = convValue
		}
		index++
	}

	// fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if isSliceSafe(arr[i]) {
			result++
		} else {
			// fmt.Println("Original Slice Not Safe")
			for j := 0; j < len(arr[i]); j++ {
				copyArr := slices.Clone(arr[i])
				if isSliceSafe(remove(copyArr, j)) {
					result++
					// fmt.Println("One Off Slice safe")
					break
				}
			}
		}
	}

	fmt.Println("Result 2 is: ", result)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isSliceSafe(in []int) bool {
	positive := false
	isNotSafe := false
	for j := 0; j < len(in)-1; j++ {
		diff_value := in[j+1] - in[j]

		if diff_value == 0 {
			isNotSafe = true
			break
		}

		if j == 0 && diff_value < 0 {
			positive = false
		} else if j == 0 && diff_value > 0 {
			positive = true
		}

		if j != 0 {
			if !positive && diff_value > 0 {
				isNotSafe = true
				break
			}

			if positive && diff_value < 0 {
				isNotSafe = true
				break
			}
		}

		if diff_value >= 4 || diff_value <= -4 {
			isNotSafe = true
			break
		}
	}
	// fmt.Println(in, cum_arr, isNotSafe)
	return !isNotSafe
}

func main() {
	part1()
	part2()
}
