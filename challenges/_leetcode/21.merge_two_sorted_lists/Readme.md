# Intuition
Basically, I thought of using a solution where the advancement of each list is linked to a principle similar to a binary tree, keeping the pointers updated for each advancement of each list and pushing the smaller values "to the left".

# Approach
I decided to solve the problem with a loop instead of a recursive function.
- The smallest `Val` between the two lists goes "to the left" of the result, that is, in the `Val` property of the `result`
- Whichever list satisfies the condition has the reference value of the pointer updated to `Next`
- If one list is smaller than the other, the values will continue to be pushed into the result until the larger list ends

# Complexity
- Time complexity: `O(n + m)` or `O(n)`
- Space complexity:
    - I think it is also `O(n)`, as the space generated depends on the number of input nodes from both lists. If I'm wrong, please correct me!

# Code
```golang []
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result

	for list1 != nil || list2 != nil {
		// To keep the values safe before advancing the pointers to the next nodes
        var l1 = list1
		var l2 = list2

    // If both values exist, compare and advance the correct list
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

    // Advance the first list only if it is larger
		if l1 != nil && l2 == nil {
			current.Next = &ListNode{
				Val: l1.Val,
			}
			list1 = list1.Next
		}

    // Advance the second list only if it is larger
		if l2 != nil && l1 == nil {
			current.Next = &ListNode{
				Val: l2.Val,
			}
			list2 = list2.Next
		}

    // Update the result pointer to the next node
		current = current.Next
	}

	return result.Next
}
```