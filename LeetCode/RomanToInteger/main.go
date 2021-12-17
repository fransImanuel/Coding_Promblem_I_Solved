package main

import "fmt"

func main() {
	fmt.Println(romanToInt("LVIII"))
}

func romanToInt(s string) int {
	mapp := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var sum int
	for i, str := range s {
		if i > 0 && mapp[string(str)] > mapp[string(s[i-1])] {
			sum = mapp[string(str)] - mapp[string(s[i-1])]
		} else {
			sum += mapp[string(str)]
		}
	}

	return sum
}
