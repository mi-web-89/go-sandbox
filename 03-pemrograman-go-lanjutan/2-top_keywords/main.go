package main

import (
	"fmt"
	"strings"
)

func GetTopKeywords(sentence string) string {
	freqs := make(map[string]int)
	var qty int = 0
	var result string

	// check length of the sentence
	tmp := strings.Fields(sentence)
	if len(tmp) == 0 {
		return ""
	}

	for _, item := range tmp {
		freqs[item]++
	}

	for key, value := range freqs {
		if value > qty || (value == qty && key < result) {
			qty = value
			result = key
		}
	}

	return result
}

func main() {
	fmt.Println("==============2. Top Keyword ===============")
	fmt.Println(GetTopKeywords("vibe coding is a coding while you are vibing along with coding")) // coding
	fmt.Println(GetTopKeywords("go makes programming fun fun"))                                   // fun
	fmt.Println(GetTopKeywords("go go go"))                                                       // go
	fmt.Println(GetTopKeywords("cats chase dogs and dogs chase dogs"))                            // dogs
	fmt.Println(GetTopKeywords("the quick brown fox"))                                            // brown
	fmt.Println(GetTopKeywords(""))

}
