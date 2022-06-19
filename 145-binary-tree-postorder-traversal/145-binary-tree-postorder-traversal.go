/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    array := make([]int, 0)
	if root != nil {
        if root.Left != nil{     
		    array = append(array, postorderTraversal(root.Left)...)
        }
        
        if root.Right != nil{
            array = append(array, postorderTraversal(root.Right)...)
        }
        
        array = append(array, root.Val)
	}
	return array
}