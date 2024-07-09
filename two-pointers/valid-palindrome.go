package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	strOnly := make([]rune,0)
	for _, char := range s { // you can loop over the string direct
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			strOnly = append(strOnly,unicode.ToLower(char)) // convert to string into small letter
		}
	}
	
	last := len(strOnly) - 1 
	for i := 0; i < len(strOnly)/2; i++ {
		if strOnly[i] != strOnly[last] {
			return false
		}
		last--
	}
	return true
}
