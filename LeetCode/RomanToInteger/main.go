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

	for i := 0; i <= len(s); {
		fmt.Println(string(s[i]))

		//Just handling if i reached over s length
		if i+1 >= len(s) {
			sum += mapp[string(s[i])]
			break
		}

		if mapp[string(s[i])] < mapp[string(s[i+1])] {
			fmt.Println("ASD")
			sum += (mapp[string(s[i+1])] - mapp[string(s[i])])
			i += 2
			continue
		} else {
			sum += mapp[string(s[i])]
			i++
			continue
		}

	}

	return sum
}
