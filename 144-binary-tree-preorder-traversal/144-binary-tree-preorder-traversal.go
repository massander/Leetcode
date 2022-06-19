/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    array := make([]int, 0)
    if root != nil{
        array = append(array, root.Val)
        left := preorderTraversal(root.Left)
        right := preorderTraversal(root.Right)
        
        array = append(array, left...)
        array = append(array, right...)
        
    }
    return array
}