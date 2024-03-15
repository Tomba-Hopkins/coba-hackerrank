package algorithm

// Question :
// Sample Input 0
// 9 - game
// 10 5 20 20 4 5 2 25 1 - skor
// Sample Output 0
// 2 4 - 2 rekor skor tinggi | 4 rekor skor rendah

// Sample Input 1
// 10
// 3 4 21 36 10 28 35 5 24 42
// Sample Output 1
// 4 0

// Answer :
func BreakingRecords(scores []int32) []int32 {
	win := 0
	lose := 0

	max := scores[0]
	min := scores[0]

	for i := 1; i < len(scores); i++{
		if scores[i] > max{
			max = scores[i]
			win += 1
		} else if scores[i] < min {
			min = scores[i]
			lose += 1
		}
	}

	return []int32{int32(win), int32(lose)}

}