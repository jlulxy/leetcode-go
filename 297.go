package main

import (
	"strconv"
	"fmt"
)

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));

const (
	PlaceHolder = "#"
	Delimiter   = "|"
)
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
func serialize(root *TreeNode) string{
	restr := ""
	//序列化
	if root == nil {
		restr += PlaceHolder
		return  restr
	}
	return  serilaizeDe(root)
}

func serilaizeDe(root *TreeNode) string{
	str := ""
	if root == nil {
		str = str + PlaceHolder + Delimiter
		return  str
	}
	str = strconv.FormatInt(int64(root.Val),10) + Delimiter
	return str + serilaizeDe(root.Left)+ serilaizeDe(root.Right)

}
func deserialize(data string) *TreeNode{
	loc := 0
	if string(data[loc]) == PlaceHolder {
		return nil
	}
	return deserilaizeDe(data,&loc)
}
func deserilaizeDe ( data string,loc *int) *TreeNode{
	lenStr := *loc
	//数组结束或遇到占位返回 nil 同时指针后移两位
	if string(data[*loc]) == PlaceHolder || *loc >= len(data){
		*loc+=2
		return  nil
	}
	//正常数据指针后移
	for  string(data[*loc]) != Delimiter  {
		*loc ++
	}
	val,_ := strconv.ParseInt(string(data[lenStr:*loc]),10,64)
	//取数后跨个分隔符
	*loc ++
	node := &TreeNode{
		Val:int(val),
		Left:deserilaizeDe(data,loc),
		Right:deserilaizeDe(data,loc),
	}
	return node
}
func main(){
	root := TreeNode{
		Val:322,
		Left:nil,
		Right:nil,
	}
	leftNode_leftNode := TreeNode{
		Val:9,
		Left:nil,
		Right:nil,

	}
	leftNode := TreeNode{
		Val:2,
		Left:&leftNode_leftNode,
		Right:nil,
	}

	rightNode:= TreeNode{
		Val:5,
		Left:nil,
		Right:nil,
	}
	root.Left = &leftNode
	root.Right = &rightNode
	str := serialize(&root)
	fmt.Println(str)
	root1 := deserialize(str)
	fmt.Println(serialize(root1))
}