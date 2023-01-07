package parser

import (
	"mckp/gordle/utils"
)

func GroupWordsByLetters(words []string) map[string][]string {
	// { [a..z]: {...words} }
	hash := make(map[string][]string)

	// For each of my words
	for i := 0; i < len(words); i++ {
		word := words[i]

		strMatchesWord := func(str string) bool {
			return str == word
		}

		// for each letter in the word
		for j := 0; j < len(word); j++ {
			char := word[j]
			key := string(char)

			// create or append
			list, exists := hash[key]
			if exists {
				// If we have already registered this word due to
				// double leters, skip adding it here
				//
				// TODO: handle duble letters
				if utils.ExistsInList(strMatchesWord, list) {
					continue
				}

				// If we do not have the word in the list already
				// we append the list and reasign the value
				hash[key] = append(list, word)
			} else {
				// If we have not seen the word yet we need to add
				// it to the list of words
				hash[key] = []string{word}
			}
		}
	}

	// We now have a hash of letters a-z and an associated
	// list of words that contain that word
	return hash
}

type LetterByPosition = map[int]map[string][]string

func GroupWordsByLetterPosition(words []string) LetterByPosition {
	// We want to take a list of words, go through each
	// and index them by each letter and its position.
	//
	// EX:
	//
	// given the words [foo, boo, bar]
	// we would produce
	/*
		{
			0: {
				b: [boo, bar],
				f: [foo] },
			},
			1: {
				o: [foo, boo],
				a: [bar]
			},
			2: {
				o: [foo, boo],
				r: [bar]
			}
		}
	*/
	// where we have all words that have b in the 0 index
	// grouped together
	hash := make(LetterByPosition)

	for i := 0; i < len(words); i++ {
		word := words[i]

		for charIndex := 0; charIndex < len(word); charIndex++ {
			char := string(word[charIndex])

			charIndexHash, charIndexExists := hash[charIndex]

			if !charIndexExists {
				charIndexHash = make(map[string][]string)
			}

			wordIndexHash, wordIndexExists := charIndexHash[char]

			if !wordIndexExists {
				wordIndexHash = []string{}
			}

			wordIndexHash = append(wordIndexHash, word)

			charIndexHash[char] = wordIndexHash

			hash[charIndex] = charIndexHash
		}
	}

	return hash
}
