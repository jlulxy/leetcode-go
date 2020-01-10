package main

func flatten(root *TreeNode)  {
	flattenRe(root)
}

func flattenRe(node *TreeNode) *TreeNode{
	if node == nil{
		return nil
	}
	nodeLeft := flattenRe(node.Left)
	nodeRight := flattenRe(node.Right)
	tmp := nodeLeft
	if nodeLeft != nil{
		node.Right = nodeLeft
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = nodeRight
	}else{
		node.Right = nodeRight
	}
	node.Left = nil
	return node
}

func flatten2 (root *TreeNode){
	if root == nil {
		return
	}
	var temp  *TreeNode
	for root != nil {
		if root.Left != nil{
			temp = root.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			temp.Right = root.Right
			root.Right = root.Left
		}
		root = root.Right
	}
	return
}