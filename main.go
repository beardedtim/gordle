package main

import (
	"fmt"
	"mckp/gordle/cli"
	"mckp/gordle/parser"
	"mckp/gordle/utils"
	"strings"
)

func main() {
	result, err := parser.ParseFile("words.txt")

	if err != nil {
		panic(err)
	}

	// knowWordsCount := len(result)

	// fmt.Printf("We have %d known words", knowWordsCount)
	// fmt.Println()
	groupedByLetterPosition := parser.GroupWordsByLetterPosition(result)
	groupedByLetter := parser.GroupWordsByLetters(result)

	// fmt.Println(groupedByLetterPosition[0]["b"])
	input := cli.InputPrompt()

	knownLetters := input.Inputs
	badLetters := input.BadLetters
	possibleWords := []string{}

	for i := 0; i < len(knownLetters); i++ {
		tuple := knownLetters[i]
		letter := tuple[0]
		isInCorrectPosition := tuple[1] == "y"

		if isInCorrectPosition {
			position := groupedByLetterPosition[i]
			wordsWithLetterInPosition := position[letter]
			possibleWords = append(possibleWords, wordsWithLetterInPosition...)
		} else {
			wordsWithLetterAtAll := groupedByLetter[letter]
			possibleWords = append(possibleWords, wordsWithLetterAtAll...)
		}
	}

	filteredWordsByPosition := []string{}

	index0Tuple := knownLetters[0]
	index1Tuple := knownLetters[1]
	index2Tuple := knownLetters[2]
	index3Tuple := knownLetters[3]
	index4Tuple := knownLetters[4]

	for _, word := range possibleWords {
		char0 := string(word[0])
		char1 := string(word[1])
		char2 := string(word[2])
		char3 := string(word[3])
		char4 := string(word[4])

		if index0Tuple[1] == "y" {
			if char0 != index0Tuple[0] {
				continue
			}
		}

		if index0Tuple[1] == "n" {
			if char0 == index0Tuple[0] {
				continue
			}
		}

		if index1Tuple[1] == "y" {
			if char1 != index1Tuple[0] {
				continue
			}
		}

		if index1Tuple[1] == "n" {
			if char1 == index1Tuple[0] {
				continue
			}
		}

		if index2Tuple[1] == "y" {
			if char2 != index2Tuple[0] {
				continue
			}
		}

		if index2Tuple[1] == "n" {
			if char2 == index2Tuple[0] {
				continue
			}
		}

		if index3Tuple[1] == "y" {
			if char3 != index3Tuple[0] {
				continue
			}
		}

		if index3Tuple[1] == "n" {
			if char3 == index3Tuple[0] {
				continue
			}
		}

		if index4Tuple[1] == "y" {
			if char4 != index4Tuple[0] {
				continue
			}
		}

		if index4Tuple[1] == "n" {
			if char4 == index4Tuple[0] {
				continue
			}
		}

		wordExists := func(check string) bool {
			return check == word
		}

		if !utils.ExistsInList(wordExists, filteredWordsByPosition) {
			filteredWordsByPosition = append(filteredWordsByPosition, word)
		}
	}

	isBadLetter := func(char string) bool {
		for _, bad := range badLetters {
			if bad == char {
				return true
			}
		}

		return false
	}

	validLetters := utils.FilterList(
		func(str string) bool {
			return str != "?"
		},
		utils.MapList(func(tuple [2]string) string {
			return tuple[0]
		}, knownLetters),
	)

	onlyValidLetters := func(word string) bool {
		for _, validLetter := range validLetters {
			index := strings.Index(word, validLetter)

			if index == -1 {
				return false
			}
		}

		return true
	}

	isValidword := func(word string) bool {
		wordCharList := strings.Split(word, "")
		contiansOnlyValidLetters := onlyValidLetters(word)

		if !contiansOnlyValidLetters {
			return false
		}

		for i := 0; i < len(wordCharList); i++ {
			char := wordCharList[i]

			if isBadLetter(char) {
				return false
			}
		}

		return true
	}

	filteredByBadLetters := utils.FilterList(isValidword, filteredWordsByPosition)
	fmt.Println(filteredByBadLetters)
}
