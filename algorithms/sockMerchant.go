package algorithm

// Question :

// Sample Input
// STDIN                       Function
// -----                       --------
// 9                           n = 9
// 10 20 20 10 10 30 50 10 20  ar = [10, 20, 20, 10, 10, 30, 50, 10, 20]
// Sample Output
// 3

// Answer :

func SockMerchant(n int32, ar []int32) int32 {

	res := 0
	
	semua := map[int32]int{}

	for i := 0; i < int(n); i++{
		semua[ar[i]]++
	}
	
	for _, a := range ar{
		if semua[a] >= 2 {
			res++
			semua[a] -= 2
		}
	}

	return int32(res)

}