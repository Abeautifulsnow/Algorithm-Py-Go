package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func inorderTraversal1(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, []int{root.Val}...)
	res = append(res, inorderTraversal1(root.Right)...)
	return res
}
