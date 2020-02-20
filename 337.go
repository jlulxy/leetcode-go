package main

import "go/ast"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

//不带记忆的递归
func robre(root *TreeNode) int {
	if root == nil {
		return 0
	}
	money := root.Val
	if root.Left!= nil{
		money += robre(root.Left.Left) + robre(root.Left.Right)

	}
	if root.Right != nil{
		money += robre(root.Right.Left) + robre(root.Right.Right)
	}
	return max337(money,robre(root.Left)+robre(root.Right))
}


func robReMem (root *TreeNode) int {
	meme := make(map[*TreeNode]int)
	return RobReMem(root,meme)

}

func RobReMem (root * TreeNode,meme map[*TreeNode]int) int {
	if root == nil{
		return 0
	}
	if val,ok := meme[root];ok{
		return val
	}
	money := root.Val
	if root.Left!= nil{
		money += RobReMem(root.Left.Left,meme) + RobReMem(root.Left.Right,meme)

	}
	if root.Right != nil{
		money += RobReMem(root.Right.Left,meme) + RobReMem(root.Right.Right,meme)
	}
	meme[root] = max337(money,RobReMem(root.Left,meme)+RobReMem(root.Right,meme))
	return meme[root]
}




func max337 (a,b int) int {
	if a >b {
		return a
	}
	return b
}
/***
//错误思路要记住 dp没有max或min就是不是全局最优
//我只返回局部的最优解或当前节点是否用到啦 就没法全局计算 因为局部最优不一定是最优的 所以两个状态拿不拿都返回
//具体选哪个决定权交给上层判断，是没办法在本层决定全局最优的
func robRe(root *TreeNode) (int,bool){
	if root == nil {
		return  0,false
	}
	right,rUse := robRe(root.Right)
	left,lUse := robRe(root.Left)
	if (rUse || lUse) == false{
		return right + left + root.Val, true
	} else if rUse && lUse{
		if root.Val > left+right{
			return root.Val,true
		}else{
			return  left +right,false
		}
	}else if rUse {
		if right  > root.Val {
			return  right+left ,false
		}else {
			return  root.Val +left,true
		}
	} else if lUse{
		if left > root.Val{
			return  left +right,false
		}else{
			return  root.Val +left,true
		}
	}
	return 0,false
}
***/

func rob (root * TreeNode) int {
	re := RobDeep(root)
	return max337(re[0],re[1])

}


func RobDeep (root *TreeNode)[]int{
	re := make([]int,2)
	if root == nil{
		return re
	}
	left := RobDeep(root.Left)
	right := RobDeep(root.Right)
	//这里要注意当前取0的时候，最优两节点也可以取0！！！
	re[0] = max337(left[0],left[1]) + max337(right[1],right[0])
	re[1] = root.Val + left[0] +right[0]
	return re
}
