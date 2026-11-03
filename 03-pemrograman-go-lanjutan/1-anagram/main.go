package main

import (
	"fmt"
	"slices"
	"unicode"
)

func CheckAnagram(word1, word2 string) bool {
	sword1 := []rune(word1) // convert to unicode character, slice of rune
	sword2 := []rune(word2)

	slices.Sort(sword1)
	slices.Sort(sword2)

	fmt.Printf("s1 : %c\n", sword1)
	fmt.Printf("s2 : %c\n", sword2)

	result := slices.EqualFunc(sword1, sword2, func(s1, s2 rune) bool {
		return unicode.ToLower(s1) == unicode.ToLower(s2)
	})

	return result
}

func main() {
	fmt.Println("==============1. Anagram ===============")
	fmt.Println(CheckAnagram("lemon", "melon"))       // true
	fmt.Println(CheckAnagram("listen", "silent"))     // true
	fmt.Println(CheckAnagram("triangle", "integral")) // true
	fmt.Println(CheckAnagram("hello", "world"))       // false
	fmt.Println(CheckAnagram("", "   "))              // false

}
