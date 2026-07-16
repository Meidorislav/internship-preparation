package main

import "fmt"

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}
 
func levelOrder(root *TreeNode) [][]int {
  var result [][]int 

	if root == nil {
		return result 
	}
	
	q := make([]*TreeNode, 0, 1)
	q = append(q, root)

	for len(q) > 0 { 
		levelSize := len(q)
		level := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left) 
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, level) 
	}
	return result
}

func main() {
	var data TreeNode 
	// input data 
	result := levelOrder(&data) 
	fmt.Println(result)
}
