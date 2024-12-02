package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1() {
	f, err := os.Open("input_test")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Split(s, "   ")
	}

	fmt.Println("Result is: ", result)
}

func part2() {
	f, err := os.Open("input_test")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		line := strings.Split(s, "   ")
	}

	fmt.Println("Result is: ", result)
}

func main() {
	part1()
	// part2()
}
