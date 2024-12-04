package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	arr := make([][]string, 0, 150)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		temp := make([]string, 0, len(line))
		for _, v := range line {
			temp = append(temp, string(v))
		}
		arr = append(arr, temp)
	}

	fmt.Println(arr)
	result = dfsHelper(arr, []string{"X", "M", "A", "S"})
	fmt.Println("Result is: ", result)
}

func dfsHelper(arr [][]string, xmas []string) int {
	result := 0

	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == xmas[0] {
				fmt.Println("Found", arr[i][j], "at", i, j)
				result += dfs(arr, i, j, xmas[1:])
			}
		}
	}
	return result
}

func dfs(arr [][]string, m, n int, xmas []string) int {
	// fmt.Println(xmas, m, n)
	// if len(xmas) == 0 {
	// 	return 1
	// }

	count := 0

	rowCheck := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colCheck := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for k := 0; k < 8; k++ {
		r := m + rowCheck[k]
		c := n + colCheck[k]
		// fmt.Println("Checking", rowCheck[k], colCheck[k])

		a := 0
		for i := 0; i < len(xmas); i++ {
			if isSafe(arr, r, c) && arr[r][c] == xmas[i] {
				fmt.Println("Found", arr[r][c], "at", r, c, rowCheck[k], colCheck[k])
				a++
				r = r + rowCheck[k]
				c = c + colCheck[k]
				fmt.Println("Going to", r, c)
			} else {
				break
			}
		}

		if a == len(xmas) {
			fmt.Println("Found XMAS")
			count++
		}
	}
	return count
}

func isSafe(arr [][]string, m, n int) bool {
	if m >= 0 && m < len(arr) &&
		n >= 0 && n < len(arr[0]) {
		return true
	}

	return false
}

func part2() {
	f, err := os.Open("input")
	result := 0
	arr := make([][]string, 0, 150)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		temp := make([]string, 0, len(line))
		for _, v := range line {
			temp = append(temp, string(v))
		}
		arr = append(arr, temp)
	}

	result = findHourGlasses(arr)

	fmt.Println("Result is: ", result)
}

func findHourGlasses(arr [][]string) int {
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			if arr[i][j] == "M" {
				fmt.Println("Found M at", i, j)
				if arr[i+1][j+1] == "A" && arr[i+2][j] == "M" && arr[i][j+2] == "S" &&
					arr[i+2][j+2] == "S" {
					fmt.Println("Found HourGlass")
					count++
				}

				if arr[i+1][j+1] == "A" && arr[i+2][j] == "S" && arr[i][j+2] == "M" &&
					arr[i+2][j+2] == "S" {
					fmt.Println("Found HourGlass")
					count++
				}
			}

			if arr[i][j] == "S" {
				fmt.Println("Found S at", i, j)
				if arr[i+1][j+1] == "A" && arr[i+2][j] == "S" && arr[i][j+2] == "M" &&
					arr[i+2][j+2] == "M" {
					fmt.Println("Found HourGlass")
					count++
				}

				if arr[i+1][j+1] == "A" && arr[i+2][j] == "M" && arr[i][j+2] == "S" &&
					arr[i+2][j+2] == "M" {
					fmt.Println("Found HourGlass")
					count++
				}

			}
		}
	}
	return count
}

func main() {
	// part1()
	part2()
}
