package algorithm

// Question :
// Function Description
// Complete the introTutorial function in the editor below. It must return an integer representing the zero-based index of .
// introTutorial has the following parameter(s):
// int arr[n]: a sorted array of integers
// int V: an integer to search for
// Returns
// int: the index of  in
// The next section describes the input format. You can often skip it, if you are using included methods or code stubs.
// Input Format
// The first line contains an integer, , a value to search for.
// The next line contains an integer, , the size of . The last line contains  space-separated integers, each a value of  where .
// The next section describes the constraints and ranges of the input. You should check this section to know the range of the input.
// Sample Input 0
// STDIN           Function
// -----           --------
// 4               V = 4
// 6               arr[] size n = 6 (not passed, see function description parameters)
// 1 4 5 7 9 12    arr = [1, 4, 5, 7, 9, 12]
// Sample Output 0
// 1
// Explanation 0
// V = 4. The value  is the  element in the array. Its index is  since the array indices start from  (see array definition under Input Format).

// Answer :

func IntroTutorial(V int32, arr []int32) int32 {
	
	for i := 0; i < len(arr); i++{
		if arr[i] == V {
			return int32(i)
		}
	}

	return 0

}