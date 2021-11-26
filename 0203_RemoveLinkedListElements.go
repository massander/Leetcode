// https://leetcode.com/problems/remove-linked-list-elements/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	start := new(ListNode)
	start.Next = head
	current := start

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}
	return start.Next
}
