package algorithm

// Question :

// Function Description
// Complete the hackerrankInString function in the editor below.
// hackerrankInString has the following parameter(s):
// string s: a string
// Returns
// string: YES or NO

// Sample Input 0 :
// hereiamstackerrank
// hackerworld
// Sample Output 0 :
// YES
// NO

// Sample Input 1 :
// hhaacckkekraraannk
// rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt
// Sample Output 1 :
// YES
// NO

// Answer :
func HackerrankInString(s string) string {

	i := 0

	main := "hackerrank"
	for _, r := range s {
		if i == len(main){
			break
		}
		if string(r) == string(main[i]) {
			i++
		}
	}

	if i == len(main) {
		return "YES"
	}

	return "NO"

}