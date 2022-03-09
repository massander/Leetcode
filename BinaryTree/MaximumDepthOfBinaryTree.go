package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	var answer int
	maxDepthHelper(root, 0, &answer)
	return answer
}

func maxDepthHelper(root *TreeNode, depth int, answer *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*answer = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}(*answer, depth)
	}

	maxDepthHelper(root.Left, depth+1, answer)
	maxDepthHelper(root.Right, depth+1, answer)
}
