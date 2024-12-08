package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type pair struct {
	first  int
	second int
}

func convertToInt(in string) int {
	digit, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	return digit
}

func part1() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	layout := make([][]rune, 0, 50)
	locations := make(map[rune][]pair)
	antiNodes := NewIntPairSet()
	i := 0

	for scanner.Scan() {
		s := scanner.Text()
		temp := make([]rune, 0, len(s))
		for j, v := range s {
			temp = append(temp, v)

			if v != '.' {
				if _, ok := locations[v]; !ok {
					locations[v] = make([]pair, 0, 5)
				}
				locations[v] = append(locations[v], pair{i, j})
			}
		}
		layout = append(layout, temp)
		i++
	}

	m := len(layout)
	n := len(layout[0])

	for _, v := range locations {
		// fmt.Println(k, v)
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				calculateAntiNodes(antiNodes, m, n, v[i], v[j])
			}
		}
	}

	// printLayout(layout)
	// fmt.Println(antiNodes.ListPairs(), antiNodes.Length())
	for _, v := range antiNodes.ListPairs() {
		layout[v.First][v.Second] = '#'
	}
	// printLayout(layout)
	fmt.Println("Result is: ", antiNodes.Length())
}

func calculateAntiNodes(antiNodes *IntPairSet, m, n int, i, j pair) {
	x_diff := j.first - i.first
	y_diff := j.second - i.second

	if isSafe(m, n, i.first-x_diff, i.second-y_diff) {
		antiNodes.AddPair(i.first-x_diff, i.second-y_diff)
	}

	if isSafe(m, n, j.first+x_diff, j.second+y_diff) {
		antiNodes.AddPair(j.first+x_diff, j.second+y_diff)
	}
}

func calculateAntiNodesResonant(antiNodes *IntPairSet, m, n int, i, j pair) {
	antiNodes.AddPair(i.first, i.second)
	antiNodes.AddPair(j.first, j.second)
	x_diff_const := j.first - i.first
	y_diff_const := j.second - i.second

	x_diff := x_diff_const
	y_diff := y_diff_const

	for isSafe(m, n, i.first-x_diff, i.second-y_diff) {
		antiNodes.AddPair(i.first-x_diff, i.second-y_diff)
		x_diff += x_diff_const
		y_diff += y_diff_const

	}

	x_diff = x_diff_const
	y_diff = y_diff_const

	for isSafe(m, n, j.first+x_diff, j.second+y_diff) {
		antiNodes.AddPair(j.first+x_diff, j.second+y_diff)
		x_diff += x_diff_const
		y_diff += y_diff_const

	}
}

func isSafe(m, n, i, j int) bool {
	if i >= m || j >= n || j < 0 || i < 0 {
		return false
	}

	return true
}

func printLayout(a [][]rune) {
	for _, v := range a {
		for _, w := range v {
			fmt.Printf("%s ", string(w))
		}
		fmt.Println()
	}

	fmt.Println("--------------")
	fmt.Println()
}

func part2() {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	layout := make([][]rune, 0, 50)
	locations := make(map[rune][]pair)
	antiNodes := NewIntPairSet()
	i := 0

	for scanner.Scan() {
		s := scanner.Text()
		temp := make([]rune, 0, len(s))
		for j, v := range s {
			temp = append(temp, v)

			if v != '.' {
				if _, ok := locations[v]; !ok {
					locations[v] = make([]pair, 0, 5)
				}
				locations[v] = append(locations[v], pair{i, j})
			}
		}
		layout = append(layout, temp)
		i++
	}

	m := len(layout)
	n := len(layout[0])

	for _, v := range locations {
		// fmt.Println(k, v)
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				calculateAntiNodesResonant(antiNodes, m, n, v[i], v[j])
			}
		}
	}

	// printLayout(layout)
	// fmt.Println(antiNodes.ListPairs(), antiNodes.Length())
	for _, v := range antiNodes.ListPairs() {
		layout[v.First][v.Second] = '#'
	}
	// printLayout(layout)
	fmt.Println("Result 2 is: ", antiNodes.Length())
}

func main() {
	part1()
	part2()
}
