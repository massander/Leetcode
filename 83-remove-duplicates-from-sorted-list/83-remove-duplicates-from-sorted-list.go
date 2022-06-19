/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }

    visited := make(map[int]bool)
    // visited[head.Val] = true
    
    start := new(ListNode); start.Next = head
    current := start
    
    for current.Next != nil{
        if _, ok := visited[current.Next.Val]; ok{
            current.Next = current.Next.Next
            continue
        }
        
        visited[current.Next.Val] = true
        current = current.Next
    }
    
    return start.Next
}