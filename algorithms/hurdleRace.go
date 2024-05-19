package algorithm

// Question :

// Sample Input 0
// k = 4
// 1 6 3 5 2
// Sample Output 0
// 2

// Answer :

func HurdleRace(k int32, height []int32) int32 {

	var max int32 = 0

	for _, r := range height{
		if r > max{
			max = r
		}
	}

	if max - k > 0 {
		return max - k
	}


	return 0

}
