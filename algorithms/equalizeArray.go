package algorithm

// Question :
// Given an array of integers, determine the minimum number of elements to delete to leave only elements of equal value.
// Function Description
// Complete the equalizeArray function in the editor below.
// equalizeArray has the following parameter(s):
// int arr[n]: an array of integers
// Returns
// int: the minimum number of deletions required

// Sample Input
// 3 3 2 1 3   arr = [3, 3, 2, 1, 3]
// Sample Output
// 2

// Answer :

func EqualizeArray(arr []int32) int32 {

	max := 0

	kardus := map[int32]int{}
	for _, r := range arr {
		kardus[r]++
	}

	for _, v := range arr{
		if kardus[v] > max {
			max = kardus[v]
		}
	}

	return int32(len(arr) - max)

}

// Awalnya gini -> :
// for _, i := range arr{
//     if kardus[i] > max {
//         max = int(i)
//     }
// }
// -> kalau begini max dibandingin sama val
// sedangkan max adalah i
// 15 max val -> sedangkan max = i -> 15 < 24