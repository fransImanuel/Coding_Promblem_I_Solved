package main

import (
	"fmt"
	"unicode"
)

func main() {
	// fmt.Println(unicode.IsLetter(90))
	var N int
	var word string
	var wordCollection []string
	var wordStringFilter []string

	//Get How Many N Element
	fmt.Scanln(&N)

	//Loop Through N to Get N Element
	for i := 0; i < N; i++ {
		fmt.Scanln(&word)
		wordCollection = append(wordCollection, word)
	}

	//This For Is For Filter And Only Get String Element
	for _, w := range wordCollection {
		// fmt.Printf("%v is %v\n", w, IsLetter(w, len(w)))
		IsLetter(w, len(w))
		wordStringFilter = append(wordStringFilter, w)
	}

	//Print Only Odd Element In Array
	fmt.Println()
	for i, v := range wordStringFilter {
		//i always start on 0 so I add 1 so it can read first element on array as 1 not 0
		if (i+1)%2 != 0 {
			fmt.Println(v)
		}
	}

}

func IsLetter(word string, length int) bool {
	var countFalse int
	// This Loop, Loop through the word and return ASCII of that word
	for _, w := range word {
		//unicode.IsLetter get ASCII and check is that ASCII stand for alphabet or not
		if !unicode.IsLetter(rune(w)) {
			countFalse++
		}
	}
	if countFalse == len(word) {
		return false
	}

	return true
}
