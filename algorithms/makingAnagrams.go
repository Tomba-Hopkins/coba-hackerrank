package algorithm

// Question :
// Sample Input
// cde
// abc
// Sample Output
// 4
// Explanation
// Delete the following characters from our two strings to turn them into anagrams:
// Remove d and e from cde to get c.
// Remove a and b from abc to get c.
// characters have to be deleted to make both strings anagrams.

// Answer :
func MakingAnagrams(s1 string, s2 string) int32 {
	abjad1 := map[string]int{}
	abjad2 := map[string]int{}


	for _, r1 := range s1 {
		abjad1[string(r1)]++
	}

	for _, r2 := range s2 {
		abjad2[string(r2)]++
	}

	res := 0

	for _, res1 := range s1 {
		if abjad1[string(res1)] > abjad2[string(res1)] {
			res++
			abjad1[string(res1)]--
		}
	}

	for _, res2 := range s2 {
		if abjad2[string(res2)] > abjad1[string(res2)] {
			res++
			abjad2[string(res2)]--
		}
	}

	return int32(res)
}