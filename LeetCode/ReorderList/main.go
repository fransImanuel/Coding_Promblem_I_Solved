package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) []int {
	var (
		i    int
		temp *ListNode
		arr  []int
	)
	newList := new(ListNode)

	length := findLength(head)

	fmt.Printf("length: %d\n", length)
	fmt.Printf("length/2: %d\n", length/2)

	for i, temp = 0, newList; i < length/2+1; i++ {
		temp = traverseAndGetValue(head, i)
		arr = append(arr, temp.Val)

		//if its odd number length
		if i >= length/2 {
			break
		}

		temp = traverseAndGetValue(head, length-i-1)
		arr = append(arr, temp.Val)
	}

	for j := 0; head != nil; j++ {
		fmt.Printf("head: %d, arr[%d]: %d\n", head.Val, j, arr[j])
		head.Val = arr[j]
		head = head.Next
		// fmt.Println(head)

	}

	fmt.Println(PrintAll(head))

	return arr

}

func findLength(head *ListNode) int {
	var (
		i    int
		temp *ListNode
	)

	for i, temp = 0, head; temp.Next != nil; i, temp = i+1, temp.Next {
	}
	return i + 1
}

func traverseAndGetValue(node *ListNode, sequence int) *ListNode {
	var (
		temp  *ListNode
		i     int
		value int
	)
	if node.Val == 0 && node.Next == nil {
		return &ListNode{}
	}

	for i, temp = 0, node; i <= sequence; i, temp = i+1, temp.Next {
		value = temp.Val
	}
	return &ListNode{Val: value}

}

func main() {

	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	// node := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val:  4,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }

	fmt.Println(reorderList(node))
}

func PrintAll(node *ListNode) []int {
	var arr []int
	temp := node
	for temp != nil {
		arr = append(arr, temp.Val)
		temp = temp.Next
	}
	return arr
}
