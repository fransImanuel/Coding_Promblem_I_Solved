package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(2))
	// fmt.Println(isPowerOfTwo(3))
	// fmt.Println(isPowerOfTwo(4))
}

func isPowerOfTwo(n int) bool {
	fmt.Println(4 & 3)
	return n > 0 && n&(n-1) == 0
}
