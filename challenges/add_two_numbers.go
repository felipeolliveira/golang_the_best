package main

import "fmt"

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 7}}}}
	l2 := ListNode{Val: 4, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 5}}}}
	// [1,9,9,7]
	// [4,9,9,5]
	// [5,8,9,3,1]

	v := addTwoNumbers(&l1, &l2)

	fmt.Printf("\nresult:")
	for v.Next != nil {
		fmt.Printf("%v", v.Val)
		v = v.Next
	}
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	currentResult := result
	rest := 0

	for l1 != nil || l2 != nil || rest != 0 {
		sum := rest

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		rest = sum / 10
		currentResult.Next = &ListNode{Val: sum % 10}
		currentResult = currentResult.Next
	}

	return result.Next
}
