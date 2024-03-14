package algorithm

import "strconv"

// Question :
// Sample Input 0
// 07:05:45PM
// Sample Output 0
// 19:05:45

// Answer :

func TimeConversion(s string) string {
	jam1, jam2 := s[0], s[1]
	jam := string(jam1) + string(jam2)

	indic1 := s[len(s) - 2]
	indicator := string(indic1)
	
	if indicator == "P" && jam != "12" {
		jamInt, _ := strconv.Atoi(jam)
		jamInt += 12
		jam = strconv.Itoa(jamInt)
	} else if indicator == "A" && jam == "12" {
		jam = "00"
	}

	result := jam + s[2:8]
	return result
}