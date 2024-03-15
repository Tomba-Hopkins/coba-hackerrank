package algorithm

import "fmt"

// Question :
// Sample Input

// 6
// Sample Output

//      #
//     ##
//    ###
//   ####
//  #####
// ######

// Answer :

func Staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		for j := n; j > i; j--{
			fmt.Printf(" ")
		}
		for k := int32(1); k <= i; k++{
			fmt.Printf("#")
		}
		fmt.Println("")
	}
}