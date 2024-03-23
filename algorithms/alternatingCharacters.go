package algorithm

// Question :
// You are given a string containing characters  and  only. Your task is to change it into a string such that there are no matching adjacent characters. To do this, you are allowed to delete zero or more characters in the string.
// Your task is to find the minimum number of required deletions.
// Sample Input :
// // 5 -> 5 test case
// AAAA
// BBBBB
// ABABABAB
// BABABA
// AAABBB
// Sample Output :
// 3
// 4
// 0
// 0
// 4

// Answer :
func AlternatingCharacters(s string) int32 {

	res := 0

	for i := 1; i < len(s); i++ {
		if string(s[i]) == string(s[i-1]) {
			res++
		}
	}

	return int32(res)

}
