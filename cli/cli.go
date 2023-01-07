package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PromptResult struct {
	Inputs     [][2]string
	BadLetters []string
}

func InputPrompt() PromptResult {
	maxLetters := 5
	strArray := make([][2]string, maxLetters)

	for i := 0; i < maxLetters; i++ {

		fmt.Printf("What is the letter for box %d? Enter \"?\" for unknown", i+1)
		fmt.Println()
		charInput := bufio.NewReader(os.Stdin)
		char, _ := charInput.ReadString('\n')

		charTrimmed := strings.TrimSpace(char)

		if charTrimmed == "?" {
			fmt.Printf("You told me you don't know %d character so we are going onto the next", i+1)
			fmt.Println()

			item := [2]string{charTrimmed, ""}
			strArray[i] = item

			continue
		}

		fmt.Print("Is this in the right spot (is it green?) Y/N")
		fmt.Println()

		isGreenInput := bufio.NewReader(os.Stdin)
		isGreen, _ := isGreenInput.ReadString('\n')

		isGreenNormalized := strings.ToLower(strings.TrimSpace(isGreen))

		var greenMessage string

		if isGreenNormalized == "y" {
			greenMessage = "was"
		} else {
			greenMessage = "was not"
		}

		fmt.Printf("You told me input %d was %s and you told me that it %s green", i+1, charTrimmed, greenMessage)
		fmt.Println()

		item := [2]string{charTrimmed, isGreenNormalized}
		strArray[i] = item
	}
	fmt.Println("What are the known \"bad\" letters (ones that are not valid for the solution)? Enter separated by a comma.")
	fmt.Println("Example: ")
	fmt.Println("a,d,t,f")

	knownBadLettersInput := bufio.NewReader(os.Stdin)
	knownBadLetters, _ := knownBadLettersInput.ReadString('\n')
	letters := strings.Split(strings.TrimSpace(knownBadLetters), ",")

	return PromptResult{
		Inputs:     strArray,
		BadLetters: letters,
	}
}
