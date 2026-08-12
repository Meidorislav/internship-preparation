/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxSoFar int) int {
    if node == nil {
        return 0
    }

    count := 0
    if node.Val >= maxSoFar {
        count = 1
        maxSoFar = node.Val
    }

    count += dfs(node.Left, maxSoFar)
    count += dfs(node.Right, maxSoFar)

    return count
}
