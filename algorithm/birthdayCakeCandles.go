package algorithm

// Question :

// Sample Input 0
// 4
// 3 2 1 3

// Sample Output 0
// 2

// Answer :

func birthdayCakeCandles(candles []int32) int32 {
    left := int32(candles[0])
    
    for i := 1; i < len(candles); i++{
        if candles[i] > left {
            left = candles[i]
        }
    }
    temp := []int32{}
    for _, candle := range candles {
        if candle == left{
            temp = append(temp, candle)
        }
    }
    result := len(temp)
    return int32(result)
}