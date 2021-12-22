package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func insertionSortList(head *ListNode) *ListNode {
// 	count := 0
// 	newNode := &ListNode{}
// 	temp := head
// 	for temp != nil {
// 		newNode = PushNode(newNode, temp.Val, count)
// 		count++
// 		temp = temp.Next
// 	}

// 	return newNode
// }

// func PushNode(currNode *ListNode, newValue, count int) *ListNode {
// 	newNode := &ListNode{Val: newValue, Next: nil}
// 	if currNode.Next == nil && currNode.Val == 0 && count == 0 {
// 		currNode = newNode
// 		return currNode
// 	}

// 	//Determine head and tail
// 	head := currNode
// 	temp := head
// 	var tail *ListNode
// 	for temp != nil {
// 		if temp.Next == nil {
// 			tail = temp
// 		}
// 		temp = temp.Next
// 	}

// 	if tail.Val < newNode.Val {
// 		tail.Next = newNode
// 		tail = newNode
// 		return head
// 	} else if head.Val >= newNode.Val {
// 		newNode.Next = head
// 		head = newNode
// 		return head
// 	} else {
// 		temp := head

// 		for temp != tail {
// 			if temp.Val <= newNode.Val && temp.Next.Val > newNode.Val {
// 				newNode.Next = temp.Next
// 				temp.Next = newNode
// 				break
// 			}
// 			temp = temp.Next
// 		}
// 	}

// 	return head

// }

func PrintAll(node *ListNode) {
	fmt.Println()
	temp := node
	for temp != nil {
		fmt.Printf("%d-> ", temp.Val)
		temp = temp.Next
	}
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	fmt.Println(dummy)
	for head != nil {
		cur := dummy
		for ; cur.Next != nil && cur.Next.Val < head.Val; cur = cur.Next {
		}
		// cur.Next, head.Next, head = head, cur.Next, head.Next

		ta := head
		tb := cur.Next
		tc := head.Next

		cur.Next = ta
		head.Next = tb
		head = tc

	}
	return dummy.Next
}

func main() {
	// a := 1
	// b := 2

	// fmt.Println(a, b)
	// a, b = b, a
	// fmt.Println(a, b)

	// a, b = b, a
	// c := 3

	// fmt.Println(a, b, c)
	// a, b, c = c, a, b
	// fmt.Println(a, b, c)

	node := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	PrintAll(insertionSortList(node))

	node = &ListNode{
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

	node = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
		},
	}

	PrintAll(insertionSortList(node))

}
