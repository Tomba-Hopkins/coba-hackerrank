package algorithm

import "strings"

// Question :
// A pangram is a string that contains every letter of the alphabet. Given a sentence determine whether it is a pangram in the English alphabet. Ignore case. Return either pangram or not pangram as appropriate.

// Sample Input 0
// We promptly judged antique ivory buckles for the next prize
// Sample Output 0
// pangram
// Sample Explanation 0
// All of the letters of the alphabet are present in the string.
// Sample Input 1
// We promptly judged antique ivory buckles for the prize
// Sample Output 1
// not pangram
// Sample Explanation 0
// The string lacks an x.

// Answer :

func Pangrams(s string) string {
	abjad := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
		"h": true,
		"i": true,
		"j": true,
		"k": true,
		"l": true,
		"m": true,
		"n": true,
		"o": true,
		"p": true,
		"q": true,
		"r": true,
		"s": true,
		"t": true,
		"u": true,
		"v": true,
		"w": true,
		"x": true,
		"y": true,
		"z": true,
	}

	last := map[string]int{}

	s = strings.ToLower(s)

	for i := 0; i < len(s); i++{
		if string(s[i]) == " "{ //kalau case ! ? sih gawat juga 
			continue
		}

		if abjad[string(s[i])]{
			last[string(s[i])]++
		}
	}

	if len(last) < len(abjad) {
		return "not pangram"
	}

	return "pangram"

}