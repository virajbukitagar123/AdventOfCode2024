package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bzick/tokenizer"
)

const (
	LParen = 1
	Mul    = 2
	RParen = 3
	Comma  = 4
	Do     = 5
	Dont   = 6
)

func readInput(path string) string {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	buffer, readErr := io.ReadAll(file)
	if readErr != nil {
		panic("Unable to read input file")
	}

	return string(buffer)
}

func DayThree() {

	garbageString := readInput("input_test")

	garbageString = strings.ReplaceAll(garbageString, "do()", "___do___")
	garbageString = strings.ReplaceAll(garbageString, "don't()", "___dont___")

	parser := tokenizer.New()
	parser.DefineTokens(LParen, []string{"("})
	parser.DefineTokens(RParen, []string{")"})
	parser.DefineTokens(Mul, []string{"mul"})
	parser.DefineTokens(Comma, []string{","})
	parser.DefineTokens(Do, []string{"___do___"})
	parser.DefineTokens(Dont, []string{"___dont___"})

	parser.AllowKeywordSymbols(tokenizer.Underscore, []rune{})

	stream := parser.ParseString(garbageString)
	defer stream.Close()

	tokenList := make([]*tokenizer.Token, 0)
	var sum int64 = 0

	for stream.IsValid() {
		token := stream.CurrentToken()
		tokenList = append(tokenList, token)

		fmt.Printf("Token Key %d Token Value: %s\n", token.Key(), token.ValueString())

		stream.GoNext()
	}

	do := true

	for i := range tokenList {
		token := tokenList[i]

		if token.Is(Do) || strings.HasSuffix(token.ValueString(), "___do___") {
			do = true
			continue
		}

		if token.Is(Dont) || strings.HasSuffix(token.ValueString(), "___dont___") {
			do = false
			continue
		}

		if !do {
			continue
		}

		if token.Is(Mul) || strings.HasSuffix(token.ValueString(), "mul") {
			lParen := tokenList[i+1]
			if !lParen.Is(LParen) {
				continue
			}

			lInteger := tokenList[i+2]
			if !lInteger.IsInteger() {
				continue
			}

			comma := tokenList[i+3]
			if !comma.Is(Comma) {
				continue
			}

			rInteger := tokenList[i+4]
			if !rInteger.IsInteger() {
				continue
			}

			rParen := tokenList[i+5]
			if !rParen.Is(RParen) {
				continue
			}

			lIntegerVal := lInteger.ValueInt64()
			rIntegerVal := rInteger.ValueInt64()

			if lIntegerVal > 999 || rIntegerVal > 999 || lIntegerVal < 0 || rIntegerVal < 0 {
				continue
			}

			fmt.Printf("Current sum:  %d  Matches:  %d %d\n", sum, lIntegerVal, rIntegerVal)

			sum = sum + (lIntegerVal * rIntegerVal)
		}
	}

	fmt.Printf("Multiplication Sum: %d", sum)
}
