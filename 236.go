/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 package main

 //初始化_ans以防未找到空指针
 var _ans *TreeNode =&TreeNode{}

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	RecursionAns(root,p,q)
  return _ans
}


//递归
//由于是二叉树，只有一个节点的子节点能包含两个节点，该节点的父节点只能有左节点有或右节点有搜索节点，
func RecursionAns(node,p,q *TreeNode) bool{
	ret := false
	if node ==nil {
		return false
	}
	mid ,left,right := 0,0,0
	if RecursionAns(node.Left,p,q) {
		left = 1
	}
	if RecursionAns(node.Right,p,q) {
		right = 1
	}
	if node == p||node == q{
		mid = 1
	}
	if (mid + left + right) > 1{
		 ret = true
	}
	//如果自己或者左右大于2证明找到该节点，由于是二叉树全局只有一个节点符合条件不知道能否体检结束
	if (mid + left +right) >= 2{
		_ans = node
	}
	return ret
}

