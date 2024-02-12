package algorithm

// Question :
// STDIN           Function
// -----           --------
// 6               arr[] size n = 6
// -4 3 -9 0 4 1   arr = [-4, 3, -9, 0, 4, 1]

// Sample output
// 0.500000 3/6
// 0.333333 2/6
// 0.166667 1/6

// Answer :

import "fmt"

func plusMinus(arr []int32) {
	n := int32(len(arr))
	negative := []int32{}
	positive := []int32{}
	zero := []int32{}

	for _, ar := range arr {
		if ar < 0 {
			negative = append(negative, ar)
		} else if ar == 0 {
			zero = append(zero, ar)
		} else {
			positive = append(positive, ar)
		}
	}

	fmt.Printf("%0.6f\n", float32(len(positive))/float32(n))
	fmt.Printf("%0.6f\n", float32(len(negative))/float32(n))
	fmt.Printf("%0.6f", float32(len(zero))/float32(n))
}