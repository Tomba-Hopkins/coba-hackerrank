package algorithm

// Question :

// Function Description
// Complete the repeatedString function in the editor below.
// repeatedString has the following parameter(s):
// s: a string to repeat
// n: the number of characters to consider
// Returns
// int: the frequency of a in the substring

// Sample Input 0
// aba
// 10
// Sample Output 0
// 7
// Sample Input 1
// a
// 1000000000000
// Sample Output 1
// 1000000000000

// Answer :

func RepeatedString(s string, n int64) (res int64) {
	
	a := 0
	
	for _, r := range s {
		if r == 'a' {
			a++
		}
	}
	fullRepeats := n / int64(len(s))
	res = fullRepeats * int64(a)

	r := n % int64(len(s))
	for i := int64(0); i < r; i++ {
		if s[i] == 'a' {
			res++
		}
	}

	return res
}

// before -> lama bjir prosesnya :
// func repeatedString(s string, n int64) (res int64) {
// 	text := ""
// 	for i := 0; i < int(n); i++{
// 		for _, r := range s {
// 			if len(text) == int(n){
// 				break
// 			}
// 			text += string(r)
// 		}
// 	}
// 	for _, t := range text{
// 		if t == 'a'{
// 			res++
// 		}
// 	}
// 	return
// }
