package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func singleNumber(nums []int) []int {
	temp := make(map[int]int)
	for _, v := range nums {
		temp[v] = temp[v] + 1
	}

	var result []int
	for i, v := range temp {
		// fmt.Printf("\ni = %d, v = %d", i, v)
		if v == 1 {
			result = append(result, i)
		}
	}

	return result
}
