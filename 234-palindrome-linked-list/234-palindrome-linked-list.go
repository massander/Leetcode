/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
    if head == nil{ return false }
    
    if head.Next == nil{ return true }
    
    firstHalfEnd := endOfFirstHalf(head)
    secondHalfStart := reverseLinkedList(firstHalfEnd.Next)
    
    result := true
    p1 := head
    p2 := secondHalfStart
    
    for (result && p2 != nil){
        if p1.Val != p2.Val{
            result = false
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    
    firstHalfEnd = reverseLinkedList(secondHalfStart)
    
    return result
}

func endOfFirstHalf(head *ListNode) *ListNode{
    if head == nil || head.Next == nil{ return head }
    
    fast := head
    slow := head
    
    for (fast.Next != nil && fast.Next.Next != nil){
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    return slow
}

func reverseLinkedList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil { return head}
    
    previousNode := new(ListNode); previousNode = nil
    currentNode := head
    
    for (currentNode != nil) {
        nextCurrent := currentNode.Next
        currentNode.Next = previousNode
        previousNode = currentNode
        currentNode = nextCurrent
    }
    
    return previousNode
}