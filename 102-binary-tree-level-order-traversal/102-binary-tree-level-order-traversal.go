/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    	var result [][]int
	var queue []*TreeNode

	queue = append(queue, root)

	for len(queue) != 0 {
		level := make([]int, 0)
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			if queue[0] != nil {
				level = append(level, queue[0].Val)
				queue = append(queue, queue[0].Left)
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		result = append(result, level)
	}

	return result[:len(result)-1]
}