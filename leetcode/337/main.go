/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    res := dfs(root)
    return max(res[0], res[1])
}

func dfs(node *TreeNode) [2]int {
    if node == nil {
        return [2]int{0, 0}
    }

    left := dfs(node.Left)
    right := dfs(node.Right)
    notRobbed := max(left[0], left[1]) + max(right[0], right[1])
    robbed := node.Val + left[0] + right[0]
    
    return [2]int{notRobbed, robbed}
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
