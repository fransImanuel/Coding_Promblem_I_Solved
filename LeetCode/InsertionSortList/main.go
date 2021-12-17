package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	newNode := &ListNode{}
	temp := head

	for temp != nil {

		//if it's first time
		if newNode.Val == 0 && newNode.Next == nil {
			newNode.Val = temp.Val
		} else {
			//if it's not the first time
			newNode = PushNode(newNode, newNode.Next, temp.Val)
		}
		// PrintAll(newNode)
		temp = temp.Next
	}

	return newNode
}

func PushNode(currNode *ListNode, nextNode *ListNode, newValue int) *ListNode {
	newNode := &ListNode{Val: newValue, Next: nil}
	PrintAll(currNode)
	if nextNode == nil && currNode.Val < newNode.Val {
		currNode.Next = newNode
		return currNode
	}

	if currNode.Val >= newNode.Val {
		newNode.Next = currNode
		currNode = newNode
	} else {
		temp := currNode

		for temp.Val < newNode.Val {
			if temp.Next.Val >= newNode.Val {
				newNode.Next = temp.Next
				temp.Next = newNode
				break
			}
			temp = temp.Next
		}

	}

	return currNode

}

func PrintAll(node *ListNode) {
	fmt.Println()
	temp := node
	for temp != nil {
		fmt.Printf("%d-> ", temp.Val)
		temp = temp.Next
	}
}

func main() {
	// node := &ListNode{
	// 	Val: 4,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 1,
	// 			Next: &ListNode{
	// 				Val:  3,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }

	node := &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}

	PrintAll(insertionSortList(node))
}
