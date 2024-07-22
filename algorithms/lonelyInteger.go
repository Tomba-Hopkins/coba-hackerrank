package algorithm

// Question :

// Given an array of integers, where all elements but one occur twice, find the unique element.
// Example
// The unique element is .
// Function Description
// Complete the lonelyinteger function in the editor below.
// lonelyinteger has the following parameter(s):
// int a[n]: an array of integers
// Returns
// int: the element that occurs only once
// Input Format
// The first line contains a single integer, , the number of integers in the array.
// The second line contains  space-separated integers that describe the values in .
// Constraints
// It is guaranteed that  is an odd number and that there is one unique element.
// , where .
// Sample Input 0
// 1
// 1
// Sample Output 0
// 1
// Explanation 0
// There is only one element in the array, thus it is unique.
// Sample Input 1
// 3
// 1 1 2
// Sample Output 1
// 2
// Explanation 1
// We have two 's, and  is unique.
// Sample Input 2
// 5
// 0 0 1 2 1
// Sample Output 2
// 2
// Explanation 2
// We have two 's, two 's, and one .  is unique.

// Answer :

func Lonelyinteger(a []int32) int32 {
	k := map[int32]int32{}
	for _, r := range a {
		k[r]++
	}
	for _, r := range a {
		if k[r] == 1  {
			return r
		}
	}
	return 0

}