// https://leetcode.com/problems/middle-of-the-linked-list/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	if head != nil {
		fast, slow := head, head

		for fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		if fast.Next != nil {
			return slow.Next
		}

		return slow
	}
	return nil
}
