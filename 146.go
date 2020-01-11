type LRUCache struct {
	dataMap map[int]*ListDataNode
	head *ListDataNode
	tail *ListDataNode
	capacity int
}
type ListDataNode struct{
	data  int
	key   int
	pre *ListDataNode
	next *ListDataNode
}


func Constructor(capacity int) LRUCache {
	head ,tail := new(ListDataNode),new(ListDataNode)
	head.next = tail
	tail.pre = head
	head.next = tail
	tail.pre = head
	Cache :=LRUCache{
		dataMap :make(map[int]*ListDataNode,capacity),
		head :head,
		tail :tail,
		capacity: capacity,
	}
	return Cache
}
func (this *LRUCache) remove( node *ListDataNode){
	node.pre.next = node.next
	node.next.pre = node.pre
    node.next = nil
    node.pre = nil
}

func (this *LRUCache) insertHead (node *ListDataNode){
	head := this.head
	node.next = head.next
	node.pre = head
	head.next.pre = node
	head.next = node
}
func (this *LRUCache) Get(key int) int {
	if node ,ok := this.dataMap[key];ok{
		this.remove(node)
		this.insertHead(node)
		return node.data
	}
	return -1
}

//把刚操作的放在最前面
func (this *LRUCache) Put(key int, value int)  {
	if  node ,ok := this.dataMap[key];ok{
		node.data = value
		this.remove(node)
		this.insertHead(node)
		return
	}
	if len(this.dataMap) >= this.capacity{
		//队列满了先去掉队列尾巴把再插入
		node := this.tail.pre
		this.remove(node)
		delete(this.dataMap,node.key)
	}
	node := &ListDataNode{
		data:value,
		key:key,
	}
	this.dataMap[key] = node
	this.insertHead(node)
}