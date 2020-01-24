package main

func findRedundantConnection(edges [][]int) []int {
	nfs := NewUnionFindSet(len(edges)+1)
	result := make([]int, 2)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		n1, n2 := nfs.getNode(a), nfs.getNode(b)
		if n1.find() == n2.find() {
			result[0],result[1] = a,b
		}else{
			nfs.union(n1,n2)
		}
	}
	return result
}

type Node struct {
	parent *Node
}
func NewNode()*Node{
	n := &Node{}
	n.parent = n
	return n
}
func (n *Node)find()*Node{
	if n == nil {
		return nil
	}
	if n == n.parent {
		return n
	}
	return n.parent.find()
}

type UnionFindSet struct {
	m []*Node
}
func NewUnionFindSet(n int) *UnionFindSet{
	ufs := &UnionFindSet{}
	ufs.m = make([]*Node,n)
	return ufs
}
func (nfs *UnionFindSet) getNode(n int) *Node {
	node := nfs.m[n]
	if node == nil {
		node = NewNode()
		nfs.m[n] = node
	}
	return node
}
func (nfs *UnionFindSet)union(a, b *Node) {
	s1, s2 := a.find(), b.find()
	if s1 == s2 {
		return
	}
	s1.parent = s2
}

