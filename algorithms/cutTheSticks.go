package algorithm

// Question :

// Function Description
// Complete the cutTheSticks function in the editor below. It should return an array of integers representing the number of sticks before each cut operation is performed.
// cutTheSticks has the following parameter(s):
// int arr[n]: the lengths of each stick
// Returns
// int[]: the number of sticks after each iteration

// Answer :
func CutTheSticks(arr []int32) (res []int32) {

	for i := 0; i < len(arr); i++{
		
		min := int32(0)
		for _, r := range arr{
			if r != 0 {
				min = r
			}
		}

		berapaKali := 0

		for j := 0; j < len(arr); j++{
			if arr[j] < min && arr[j] != 0{
				min = arr[j]
			}
		}

		for j := 0; j < len(arr); j++{

			if arr[j] <= 0{
				continue
			}
			arr[j] -= min
			berapaKali++
		}

		if berapaKali <= 0 {
			continue
		}

		res = append(res, int32(berapaKali))
		
	}

	return
}