package main

import (
	"fmt"
)

func main() {
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	// indexMap := make(map[int]int)
	// indexMap[0] = 1
	// indexMap[1] = 2

	// value, ok := indexMap[0]
	// fmt.Println(indexMap)
	// fmt.Println(value)
	// fmt.Println(ok)
}

//Brute force
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {

// 			// you can't use same element twice
// 			// then if index is same then just continue
// 			if i == j {
// 				continue
// 			}
// 			fmt.Printf("\ni = %d || j = %d", i, j)
// 			if nums[i]+nums[j] == target {
// 				fmt.Printf("MSK")
// 				return []int{i, j}
// 			}
// 		}
// 	}

// 	return []int{}
// }

//using map
func twoSum(nums []int, target int) []int {
	mapp := make(map[int]int)
	// for i, m := range nums {

	// 	//I cant use library like math.Abs() T.T
	// 	mapp[target-m] = i

	// }

	//bayangkan target - n = index
	//n = target + index
	//kalau misalnya target-n sudah ada nilainya itu artinya
	for j, n := range nums {
		if value, ok := mapp[target-n]; ok {
			return []int{value, j}
		}
		mapp[n] = j
	}

	return []int{}

}

// func twoSum(nums []int, target int) []int {
// 	indexMap := make(map[int]int)
// 	for currIndex, currNum := range nums {
// 		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
// 			return []int{requiredIdx, currIndex}
// 		}
// 		indexMap[currNum] = currIndex
// 	}
// 	return []int{}
// }
