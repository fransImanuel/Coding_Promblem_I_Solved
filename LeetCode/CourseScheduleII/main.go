package main

import "fmt"

func main() {
	fmt.Println(findOrder(4, [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2}}))

}

func findOrder(numCourses int, pr [][]int) []int {

	isVisited, stack := make([]bool, numCourses), make([]int, numCourses)
	adjacentLists := make([][]int, numCourses)
	// adjacentList[0] = []int{1, 2}

	for _, p := range pr {
		adjacentLists[p[1]] = append(adjacentLists[p[1]], p[0])
	}

	fmt.Println(adjacentLists)

	for i, adjacentList := range adjacentLists {
		dfs(i, adjacentList, isVisited, stack)

	}

	return pr[0]
}

func dfs(vertices int, adjacentArray []int, isVisited []bool, stack []int) {

	if isVisited[vertices] {
		return
	} else {
		isVisited[vertices] = true
	}

	for i, adjacent := range adjacentArray {
		fmt.Println(adjacent, i)
		dfs(adjacent, adjacentArray, isVisited, stack)
	}

}

// func dfs(visited *[]bool, stack, storeValue *[]int, currentVertex int, pr [][]int) {
// 	for i, array := range pr {
// 		for j, arrayValue := range array {
// 			if arrayValue == currentVertex && j == 1 {
// 				dfs(visited, stack, storeValue, array[j-1], pr)
// 			}
// 		}
// 	}
// }

//1. create relation according to input array
//2. create stack for save least visited node
