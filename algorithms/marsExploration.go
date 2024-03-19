package algorithm

// Question :

// Sample Input 0
// SOSSPSSQSSOR
// Sample Output 0
// 3
// Explanation 0
//  = SOSSPSSQSSOR, and signal length . They sent  SOS messages (i.e.: ).
// Expected signal: SOSSOSSOSSOS
// Recieved signal: SOSSPSSQSSOR
// Difference:          X  X   X
// Sample Input 1
// SOSSOT
// Sample Output 1
// 1
// Explanation 1
//  = SOSSOT, and signal length . They sent  SOS messages (i.e.: ).
// Expected Signal: SOSSOS
// Received Signal: SOSSOT
// Difference:           X
// Sample Input 2
// SOSSOSSOS
// Sample Output 2
// 0
// Explanation 2
// Since no character is altered, return 0.

// Answer :

func MrsExploration(s string) int32 {
	res := 0

	for i := 0; i < len(s); i += 3 {
		if string(s[i]) != "S"{
			res++
		}
		if string(s[i + 1]) != "O"{
			res++
		}
		if string(s[i + 2]) != "S"{
			res++
		}
	}

	return int32(res)

}