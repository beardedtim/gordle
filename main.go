package main

import (
	"fmt"
	"mckp/gordle/cli"
	"mckp/gordle/parser"
)

func main() {
	result, err := parser.ParseFile("words.txt")

	if err != nil {
		panic(err)
	}

	knowWordsCount := len(result)

	fmt.Printf("We have %d known words", knowWordsCount)
	fmt.Println()

	// groupedByLetters := parser.GroupWordsByLetters(result)
	// groupedByLetterPosition := parser.GroupWordsByLetterPosition(result)
	// fmt.Println(groupedByLetterPosition[0]["b"])
	input := cli.InputPrompt()

	fmt.Println(input)
}
