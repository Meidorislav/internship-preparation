/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getSumOfGreaterNodes(node *TreeNode, sum *int) {
    if node == nil {
        return
    }
    getSumOfGreaterNodes(node.Right, sum)

    *sum += node.Val
    node.Val = *sum

    getSumOfGreaterNodes(node.Left, sum)
}

func bstToGst(root *TreeNode) *TreeNode {
    sum := 0
    getSumOfGreaterNodes(root, &sum)
    return root
}
