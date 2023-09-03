package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 一路沿着最左侧或者最右侧递归查找下去, 深度检索
// 有值加1, 无值返回0
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 宽度检索
func maxDepthWithBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)

	for len(nodeList) > 0 {
		result++

		n := len(nodeList)
		for a := 0; a < n; a++ {
			if nodeList[0].Left != nil {
				nodeList = append(nodeList, nodeList[0].Left)
			}
			if nodeList[0].Right != nil {
				nodeList = append(nodeList, nodeList[0].Right)
			}
			nodeList = nodeList[1:]
		}
	}

	return result
}
