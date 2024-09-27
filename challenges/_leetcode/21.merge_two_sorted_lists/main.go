// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// Example 2:
// Input: list1 = [], list2 = []
// Output: []

// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]

// Constraints:
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

package main

func main() {
	l1 := &ListNode{Val: 5}
	l1.Next = &ListNode{Val: 10}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 2}
	l2.Next.Next = &ListNode{Val: 4}

	r := mergeTwoLists(l1, l2)

	for r != nil {
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil || list2 != nil {
		var l1 = list1
		var l2 = list2

		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				current.Next = &ListNode{
					Val: l1.Val,
				}
				list1 = list1.Next
			} else {
				current.Next = &ListNode{
					Val: l2.Val,
				}
				list2 = list2.Next
			}
		}

		if l1 != nil && l2 == nil {
			current.Next = &ListNode{
				Val: l1.Val,
			}
			list1 = list1.Next
		}

		if l2 != nil && l1 == nil {
			current.Next = &ListNode{
				Val: l2.Val,
			}
			list2 = list2.Next
		}

		current = current.Next
	}

	return result.Next
}
