package algorithm

// Question :
// Sample Input
// middle-Outz
// 2
// Sample Output
// okffng-Qwvb
// Explanation
// Original alphabet:      abcdefghijklmnopqrstuvwxyz
// Alphabet rotated +2:    cdefghijklmnopqrstuvwxyzab
// m -> o
// i -> k
// d -> f
// d -> f
// l -> n
// e -> g
// -    -
// O -> Q
// u -> w
// t -> v
// z -> b

// Answer :
func CaesarCipher(s string, k int32) string {

	res := ""
	
	abjad := map[string]int32{}
	abjad2 := map[int32]string{}

	hehe := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i, h := range hehe {
		abjad[string(h)] = int32(i + 1)
	}

	for i, h := range hehe{
		abjad2[int32(i + 1)] = string(h)
	}
	
	for i := 0; i < len(s); i++{
		str := string(s[i])
		strInt := abjad[str]

		if strInt == 0 {
			res += str
		}
		
		if strInt + k > 52 {
			strInt += k - 52
			} else {
				strInt += k
			}
			
			newStr := abjad2[strInt]
			
		res += newStr

	}


	return res
}

// main function dari sono nya error gajelas jir