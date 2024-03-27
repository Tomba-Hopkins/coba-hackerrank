package mathematics

// Question :

// Function Description
// Complete the maximumDraws function in the editor below.
// maximumDraws has the following parameter:
// int n: the number of colors of socks
// Returns
// int: the minimum number of socks to remove to guarantee a matching pair.
// Input Format
// The first line contains the number of test cases, .
// Each of the following  lines contains an integer .
// Sample Input
// 2
// 1
// 2
// Sample Output
// 2
// 3
// Explanation
// Case 1 : Only 1 color of sock is in the drawer. Any  will match.
// Case 2 : 2 colors of socks are in the drawer. The first two removed may not match. At least  socks need to be removed to guarantee success.

// Answer :

func MaximumDraws(n int32) int32 {
    return n + 1

}

