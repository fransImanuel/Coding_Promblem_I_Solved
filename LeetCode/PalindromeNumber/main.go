package main

import "fmt"

func main() {

	fmt.Println(isPalindrome(121))
}

// func isPalindrome(x int) bool {
//     var newNumber int = 0
//     n := x
//     for n > 0 {
//         lastDigit := n % 10
//         newNumber *= 10
//         newNumber += lastDigit
//         n /= 10
//     }

//     if newNumber == x {return true}
//     return false
// }

func isPalindrome(x int) bool {
	if x >= 0 && x == Decode(x, 0) {
		return true
	}
	return false
}

func Decode(x, s int) int {
	if x < 10 {
		return s*10 + x
	} else {
		temp := x % 10
		return Decode(x/10, s*10+temp)
	}
}
