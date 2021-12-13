package main

import "fmt"

func main() {

	// arr := [3]int{1, 2, 3}
	// fmt.Println(string(93))
	fmt.Println(isValid("()"))     //true
	fmt.Println(isValid("()[]{}")) //true
	fmt.Println(isValid("(]"))     //false
	fmt.Println(isValid("([)]"))   //false
	fmt.Println(isValid("{[]}"))   //true
	fmt.Println(isValid("([}}])")) //false
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var arr []rune

	for _, str := range s {
		// ( = 40
		// ) = 41
		// [ = 91
		// ] = 93
		// { = 123
		// } = 125
		if str == 40 || str == 91 || str == 123 {
			arr = append(arr, str)
		} else if str == 41 && len(arr) != 0 && arr[len(arr)-1] == 40 {
			arr = arr[:len(arr)-1]
		} else if str == 93 && len(arr) != 0 && arr[len(arr)-1] == 91 {
			arr = arr[:len(arr)-1]
		} else if str == 125 && len(arr) != 0 && arr[len(arr)-1] == 123 {
			arr = arr[:len(arr)-1]
		} else {
			return false
		}
	}

	fmt.Println(arr)

	return len(arr) == 0
}
