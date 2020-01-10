package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	//二叉树层次便利，返回每层最右边的
	ret := make([]int,0)
	if root == nil {
		return ret
	}
	//lastLevelList
	lastLevelList := make([]*TreeNode,0,100)
	lastLevelList = append(lastLevelList,root)
	//保证可弹出下一层
	for len(lastLevelList) != 0 {
		//addlastLevel lastnode
		ret = append(ret,lastLevelList[len(lastLevelList)-1].Val)
		nowLevelList := make([]*TreeNode,0,100)
		//便利当前层把存储下一层
		for i := range lastLevelList{
			if lastLevelList[i].Left !=nil{
				nowLevelList = append(nowLevelList,lastLevelList[i].Left)
			}
			if lastLevelList[i].Right != nil{
				nowLevelList = append(nowLevelList,lastLevelList[i].Right)
			}
		}
		//当前层变成下一层
		lastLevelList = nowLevelList
	}
	return ret
}
