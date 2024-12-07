package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func convertToInt(in string) int {
	digit, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}

	return digit
}

func part1() {
	f, err := os.Open("input")
	result := 0

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	g := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())
	orders := make([][]int, 0, 100)
	line := 0

	for scanner.Scan() {
		s := scanner.Text()
		if !strings.Contains(s, "|") {
			break
		}
		nodes := strings.Split(s, "|")
			
		}
	}

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			continue
		}

		orders = append(orders, make([]int, 0, 30))
		pages := strings.Split(s, ",")
		for _, v := range pages {
			orders[line] = append(orders[line], convertToInt(v))
		}
		line++
	}

	topSort, err := graph.TopologicalSort(g)
	if err != nil {
		log.Fatal(err)
	}

	topSortMap := make(map[int]int)
	for i, v := range topSort {
		topSortMap[v] = i
	}

	fmt.Println("Map is ", topSortMap)
	fmt.Println()
	for _, v := range orders {
		// isCorrect := false

		indexArr := make([]int, 0, len(v))
		for _, k := range v {
			index, ok := topSortMap[k]
			if ok {
				indexArr = append(indexArr, index)
			}
		}

		fmt.Println(v, indexArr)
		if isCorrectOrder(indexArr) {
			fmt.Println("Order Correct")
			result += v[(len(v) / 2)]
		}
	}

	fmt.Println("Result is: ", result)
}

func isCorrectOrder(indexArr []int) bool {
	for i := 0; i < len(indexArr)-1; i++ {
		if indexArr[i] > indexArr[i+1] {
			return false
		}
	}
	return true
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

// func main() {
// 	part1()
// 	// part2()
// }
