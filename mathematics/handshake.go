package mathematics

// Question :

// Example
// There are  attendees, ,  and .  shakes hands with  and , and  shakes hands with . Now they have all shaken hands after  handshakes.
// Function Description
// Complete the handshakes function in the editor below.
// handshakes has the following parameter:
// int n: the number of attendees
// Returns
// int: the number of handshakes
// Input Format
// The first line contains the number of test cases .
// Each of the following  lines contains an integer, .
// Sample Input
// 2
// 1
// 2
// Sample Output
// 0
// 1
// Explanation
// Case 1 : The lonely board member shakes no hands, hence 0.
// Case 2 : There are 2 board members, so 1 handshake takes place.

// Answer :
func Handshake(n int32) int32 {  
    return ((n - 1) * n) / 2

}