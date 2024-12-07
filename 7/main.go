package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToInt64(in string) uint64 {
	digit, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		panic(err)
	}

	return digit
}

func convertToInt(in string) int {
	digit, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return digit
}

func recursion(num uint64, list []int, currValue uint64) bool {
	if currValue > num {
		return false
	}

	if len(list) == 0 {
		return currValue == num
	}

	return recursion(num, list[1:], uint64(list[0])*currValue) || recursion(num, list[1:], uint64(list[0])+currValue)
}

func isEquationCorrect(num uint64, list []int) bool {
	return recursion(num, list[1:], uint64(list[0]))
}

// -------------------

func recursionWithConcat(num uint64, list []int, currValue uint64) bool {
	if currValue > num {
		return false
	}

	if len(list) == 0 {
		if currValue == num {
			fmt.Println(currValue)
		}
		return currValue == num
	}

	// fmt.Println(currValue, list[0], convertToInt64(fmt.Sprintf("%v%v", currValue, list[0])))

	return recursionWithConcat(num, list[1:], uint64(list[0])*currValue) || recursionWithConcat(num, list[1:], uint64(list[0])+currValue) || recursionWithConcat(num, list[1:],
		convertToInt64(fmt.Sprintf("%v%v", currValue, list[0])))
}

func isEquationCorrectWithConcat(num uint64, list []int) bool {
	return recursionWithConcat(num, list[1:], uint64(list[0]))
}

func part2() {
	f, err := os.Open("input")
	var result uint64 = 0

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		list := strings.Split(s, ":")
		numsList := strings.Split(list[1], " ")
		numsList = numsList[1:]
		nums := make([]int, 0, len(numsList))
		for _, v := range numsList {
			nums = append(nums, convertToInt(v))
		}

		if isEquationCorrectWithConcat(convertToInt64(list[0]), nums) {
			result += convertToInt64(list[0])
		}
	}

	fmt.Println("Result 2 is: ", result)
}

func main() {
	//part1()
	part2()
}

func part1() {
	f, err := os.Open("input")
	var result uint64 = 0

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		list := strings.Split(s, ":")
		numsList := strings.Split(list[1], " ")
		numsList = numsList[1:]
		nums := make([]int, 0, len(numsList))
		for _, v := range numsList {
			nums = append(nums, convertToInt(v))
		}

		if isEquationCorrect(convertToInt64(list[0]), nums) {
			result += convertToInt64(list[0])
		}
	}

	fmt.Println("Result is: ", result)
}
