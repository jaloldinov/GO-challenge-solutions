package gochallenges

// first solution
func Reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// second solution
func Reverse(input string) string {
	reversed := ""
	for _, c := range input {
		reversed = string(c) + reversed
	}
	return reversed
}

// third looks like second one
func Reverse(input string) (reversed string) {
	for _, v := range input {
		reversed = string(v) + reversed
	}
	return reversed
}

// fourth solution
func reverseString(str string) string {
	runes := []rune(str)
	reversed := make([]rune, len(runes))

	for i, j := len(runes)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = runes[i]
	}

	return string(reversed)
}
