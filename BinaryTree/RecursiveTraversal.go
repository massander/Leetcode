package main

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
	if root != nil {
		array = append(array, root.Val)
		left := preorderTraversal(root.Left)
		right := preorderTraversal(root.Right)

		array = append(array, left...)
		array = append(array, right...)

	}
	return array
}

func inorderTraversal(root *TreeNode) []int {
    array := make([]int, 0)
	if root != nil {
        if root.Left != nil{     
		    array = append(array, inorderTraversal(root.Left)...)
        }

        array = append(array, root.Val)
        
        if root.Right != nil{
            array = append(array, inorderTraversal(root.Right)...)
        }
	}
	return array
}