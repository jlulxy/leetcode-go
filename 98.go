package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root ==nil{
		return true
	}
	re,_,_:=isValidBSTRe(root)
	return re
}

// isValidBSTLeft 左边返回最大值
// 判断思路返回字数的左最大和右再和跟比较判断条件比较复杂
// 收集意见的方式需要大量的记录和调查比较
func isValidBSTRe (root * TreeNode) (bool,int,int){
	// 为假的情况先返回
	// 左边大于返回false
	var min,max,leftMax,rightMin int
	var left,right ,leftNil,rightNil bool
	if root.Left != nil && root.Left.Val > root.Val{
		return false,0,0
	}
	//右边小返回false
	if root.Right != nil && root.Right.Val < root.Val{
		return false,0,0
	}
	if root.Right == nil && root.Left ==nil{
		return true,root.Val,root.Val
	}
	// 判断左情况
	if root.Left == nil {
		min = root.Val
		left = true
		leftNil = true

	}else{
		left ,min,leftMax = isValidBSTRe(root.Left)
	}
	if root.Right == nil {
		max = root.Val
		rightNil = true
		right = true
	} else {
		right,rightMin,max = isValidBSTRe(root.Right)
	}
	if left && right && (leftMax < root.Val || leftNil) && (rightMin > root.Val || rightNil ){
		return true,min,max
	}
	return false,0,0
}


// _isValidBST从上到到下限制子树约束条件，简单方便
// 立规矩的方式，你就按照我的范围搞，符合条件你愿意是啥是啥！
func _isValidBST(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || root.Val >= max {
		return false
	}

	return _isValidBST(root.Left, min, root.Val) && _isValidBST(root.Right, root.Val, max)

}



