package algorithm

// Question :

// Sample Input
// saveChangesInTheEditor
// Sample Output
// 5

// Answer :

func Camelcase(s string) int32 {
	res := 0
	abjad := map[string]int{
		"A": 1,
		"B": 1,
		"C": 1,
		"D": 1,
		"E": 1,
		"F": 1,
		"G": 1,
		"H": 1,
		"I": 1,
		"J": 1,
		"K": 1,
		"L": 1,
		"M": 1,
		"N": 1,
		"O": 1,
		"P": 1,
		"Q": 1,
		"R": 1,
		"S": 1,
		"T": 1,
		"U": 1,
		"V": 1,
		"W": 1,
		"X": 1,
		"Y": 1,
		"Z": 1,
	}

	for _, r := range s {
		if abjad[string(r)] > 0 {
			res++
		}
	}
	return int32(res + 1)

}


// func camelcase(s string) int32 {
// 	res := 0
// 	for _, r := range s {
// 		if string(r) == strings.ToUpper(string(r)) {
// 			res++
// 		} 
// 	}
// 	return int32(res + 1)
	
// }