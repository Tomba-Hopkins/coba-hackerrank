package algorithm

// Question :
// Function Description
// Complete the stringConstruction function in the editor below. It should return the minimum cost of copying a string.
// stringConstruction has the following parameter(s):
// s: a string

// Sample Input
// abcd
// abab
// Sample Output
// 4
// 2

// Answer :
func stringConstruction(s string) int32 {
	abjad := map[rune]int{}

	for _, r := range s {
		abjad[r]++
	}

	res := int32(len(abjad))

	return res

}