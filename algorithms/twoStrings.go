package algorithm

// Question :

// Sample Input
// hello
// world
// = YES
// hi
// world
// = NO
// Sample Output
// YES
// NO

// Answer :

func TwoStrings(s1 string, s2 string) string {
	abjad1 := map[rune]bool{}
	abjad2 := map[rune]bool{}

	for _, r1 := range s1 {
		abjad1[r1] = true
	}

	for _, r2 := range s2 {
		abjad2[r2] = true
	}

	for _, x := range s1 {
		if abjad1[x] && abjad2[x] {
			return "YES"
		}
	}

	return "NO"

}
