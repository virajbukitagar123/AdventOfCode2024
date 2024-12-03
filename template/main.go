package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1() {
	f, err := os.Open("input_test")
	result := 0

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
	}

	fmt.Println("Result is: ", result)
}

// func part2() {
// 	f, err := os.Open("input_test")
//  result := 0

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		s := scanner.Text()
// 	}

// 	fmt.Println("Result is: ", result)
// }

func main() {
	part1()
	// part2()
}
