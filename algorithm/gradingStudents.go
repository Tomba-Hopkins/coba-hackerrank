package algorithm

import "math"

// Question :
// Sample Input 0
// 4
// 73
// 67
// 38
// 33
// Sample Output 0
// 75
// 67
// 40
// 33

// Answer :
func gradingStudents(grades []int32) []int32 {
    result := []int32{}
    
    for _, grade := range grades {
        flot := float64(grade)
        flot /= 5.0
        naik := int32(math.Ceil(flot) * 5)
        
        if naik - grade < 3 && naik >= 40 {
            result = append(result, naik)
        } else {
            result = append(result, grade)
        }
    }
    return result

}