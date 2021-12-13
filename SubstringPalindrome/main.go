package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		str    string
		start  int
		end    int
		len1   int
		len2   int
		length float64
	)
	fmt.Scanln(&str)

	/*
		this loop is looping through character and define some middle
		e. g abba string abb | bba or there is a case that the sring is even number like racecar
		so both of this case represent by len1 and len2 because I dont know what gonna hit the case

		start and end variable is used to deterimne substring index from start to end
	*/
	for i := 0; i < len(str); i++ {
		len1 = ExpandFromMiddle(str, i, i)
		len2 = ExpandFromMiddle(str, i, i+1)
		length = math.Max(float64(len1), float64(len2))
		if int(length) > end-start {
			start = i - ((int(length) - 1) / 2)
			end = i + (int(length) / 2)
		}
	}

	fmt.Println(str[start : end+1])
}

func ExpandFromMiddle(s string, left, right int) int {
	if len(s) == 0 || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

/*
	@Logic itself is true but not acceptable in kattis
	I learn this by watching algorithm explaination on youtube  https://www.youtube.com/watch?v=y2BD4MJqV20
*/
