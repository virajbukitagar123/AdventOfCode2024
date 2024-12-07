package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type puzzle struct {
	layout [][]string
	guard  []int
}

type pair struct {
	p []int
}

func (p puzzle) printPuz() {
	for _, v := range p.layout {
		for _, w := range v {
			fmt.Printf("%s ", w)
		}
		fmt.Println()
	}

	fmt.Println("--------------")
	fmt.Println()
}

func (p *puzzle) traverseWithLoop(start [][]string) {
	loopCount := 0
	for m := range p.layout {
		for n := range p.layout[m] {
			// fmt.Println(m, n)
			p.layout = clonePuz(start)
			fmt.Println(p.layout)
			if p.layout[m][n] == "#" {
				continue
			}
			p.layout[m][n] = "#"
			loopFound := false

			for !p.isGuardEdged() && !loopFound {
				a, b := p.getGuardPosition()

				if p.layout[a][b] == "^" {
					for !p.isGuardEdged() && a > 0 && p.layout[a-1][b] != "#" && !loopFound {
						if p.layout[a-1][b] == ">" {
							loopCount++
							loopFound = true
							break
						}
						a--
						p.layout[a][b] = "^"
						p.setGuardPosition(a, b)

					}

					if !p.isGuardEdged() && p.layout[a-1][b] == "#" {
						p.turnPosition()
						// p.printPuz()
					}
				}

				if p.layout[a][b] == ">" {
					for !p.isGuardEdged() && b < len(p.layout[a]) && p.layout[a][b+1] != "#" && !loopFound {
						if p.layout[a][b+1] == "V" {
							loopCount++
							loopFound = true
							break
						}
						b++
						p.layout[a][b] = ">"
						p.setGuardPosition(a, b)
					}

					if !p.isGuardEdged() && p.layout[a][b+1] == "#" {
						p.turnPosition()
						// p.printPuz()
					}
				}

				if p.layout[a][b] == "V" {
					for !p.isGuardEdged() && a < len(p.layout) && p.layout[a+1][b] != "#" && !loopFound {
						if p.layout[a+1][b] == "<" {
							loopCount++
							loopFound = true
							break
						}
						a++
						p.layout[a][b] = "V"
						p.setGuardPosition(a, b)
					}

					if !p.isGuardEdged() && p.layout[a+1][b] == "#" {
						p.turnPosition()
						// p.printPuz()
					}
				}

				if p.layout[a][b] == "<" {
					for !p.isGuardEdged() && b > 0 && p.layout[a][b-1] != "#" && !loopFound {
						if p.layout[a][b-1] == "^" {
							loopCount++
							loopFound = true
							break
						}
						b--
						p.layout[a][b] = "<"
						p.setGuardPosition(a, b)
					}

					if !p.isGuardEdged() && p.layout[a][b-1] == "#" {
						p.turnPosition()
						// p.printPuz()
					}
				}
			}

			p.printPuz()
			fmt.Println(loopFound)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

		}
	}
	fmt.Println("Loop Count:", loopCount)
}

func (p puzzle) isGuardEdged() bool {
	a, b := p.getGuardPosition()

	if a == 0 || b == 0 || a == len(p.layout)-1 || b == len(p.layout[a])-1 {
		return true
	}

	return false
}

func (p *puzzle) turnPosition() {
	a, b := p.getGuardPosition()
	switch p.layout[a][b] {
	case "^":
		p.layout[a][b] = ">"
	case ">":
		p.layout[a][b] = "V"
	case "V":
		p.layout[a][b] = "<"
	case "<":
		p.layout[a][b] = "^"
	}
}

func (p puzzle) getNextTurnPos() (int, int) {
	a, b := p.getGuardPosition()

	switch p.layout[a][b] {
	case "^":
		return a, b + 1
	case ">":
		return a + 1, b
	case "V":
		return a, b - 1
	case "<":
		return a - 1, b
	}
	return -1, -1
}

func (p puzzle) getGuardPosition() (int, int) {
	return p.guard[0], p.guard[1]
}

func (p *puzzle) setGuardPosition(a, b int) {
	p.guard[0] = a
	p.guard[1] = b
}

func (p *puzzle) initializeGuardPosition() {
	for i, v := range p.layout {
		for j, w := range v {
			if w == "^" || w == "<" || w == "V" || w == ">" {
				p.guard[0] = i
				p.guard[1] = j
			}
		}
	}
}

func (p puzzle) countResult() int {
	count := 0

	for _, arr := range p.layout {
		for _, v := range arr {
			if v == "X" {
				count++
			}
		}
	}
	return count
}

func part1() {
	f, err := os.Open("input_test")
	result := 0

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var puz puzzle = puzzle{make([][]string, 0, 130), make([]int, 2)}
	var clone [][]string

	for scanner.Scan() {
		s := scanner.Text()
		temp := make([]string, 0, len(s))
		temp2 := make([]string, 0, len(s))
		temp = append(temp, strings.Split(s, "")...)
		temp2 = append(temp2, strings.Split(s, "")...)
		puz.layout = append(puz.layout, temp)
		clone = append(clone, temp2)
	}

	puz.printPuz()
	puz.initializeGuardPosition()
	puz.traverseWithLoop(clone)
	puz.printPuz()
	result = puz.countResult()
	fmt.Println("Result is: ", result+1)
}

func main() {
	part1()
	// part2()
}

func clonePuz(room [][]string) [][]string {
	cln := make([][]string, len(room))
	m := len(room[0])
	for i, row := range room {
		cln[i] = make([]string, m)
		copy(cln[i], row)
	}
	return cln
}
