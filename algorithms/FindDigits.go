package algorithm

import "strconv"

// Question :

// Function Description
// Complete the findDigits function in the editor below.
// findDigits has the following parameter(s):
// int n: the value to analyze
// Sample Input
// 12
// 1012
// Sample Output
// 2
// 3

// Answer :
func FindDigits(n int32) (res int32) {

	str := strconv.Itoa(int(n))

	for _, s := range str{
		intS, _ := strconv.Atoi(string(s))
		if s == '0'{
			continue
		}

		if n % int32(intS) == 0 {
			res++
		}
	}

	return

}