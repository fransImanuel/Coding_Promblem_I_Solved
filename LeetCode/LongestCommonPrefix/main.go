package main

import "fmt"

func main() {
	original := "cat"
	err := original[3]

	fmt.Println(err)

}

func longestCommonPrefix(strs []string) string {
	var prefix []rune
	for i, str := range strs {
		for j, s := range str {
			if i == 0 {
				prefix = append(prefix, s)
			} else {
				if j > len(str)-1 || j > len(prefix)-1 {
					break
				}
				if prefix[j] != s {
					prefix = prefix[:j]
				} else if len(prefix) > len(str) && prefix[j-1] == rune(str[j-1]) {
					prefix = prefix[:len(str)-1]
				}
			}
		}
	}

	return string(prefix)
}
