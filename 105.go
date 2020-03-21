package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildfTreeRe(preorder, inorder)
}

// 递归从先跟中找到跟，从中跟中找到位置 ，划分左右子树
func buildfTreeRe(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	tmp := &TreeNode{
		Val: preorder[0],
	}
	k := findK(preorder[0], inorder)
	tmp.Left = buildfTreeRe(preorder[1:k+1], inorder[0:k])
	tmp.Right = buildfTreeRe(preorder[k+1:], inorder[k+1:])
	return tmp
}

func findK(find int, inorder []int) int {
	for k, v := range inorder {
		if v == find {
			return k
		}
	}
	return -1
}
