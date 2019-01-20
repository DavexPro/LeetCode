type Node struct {
	Prev *Node
	Next *Node
	Key int
	Val int
}

type LRUCache struct {
	Capacity int
	Head *Node
	Tail *Node
	Store map[int]*Node
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache{Capacity:capacity, Store: make(map[int]*Node)}
	return lru
}


func (c *LRUCache) Get(key int) int {
	node, ok := c.Store[key]
	if ok {
		c.removeNode(node)
		c.appendToHead(node)
		return node.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int)  {
	node, ok := c.Store[key]

	// if out of capacity, eject the least used one
	if !ok && len(c.Store) >= c.Capacity {
		last := c.Tail
		if last != nil {
			delete(c.Store, last.Key)
			c.removeNode(last)
		}
	}

	if ok {
		c.removeNode(node)
	}

	node = &Node{Key:key, Val:value}
	c.Store[key] = node
	c.appendToHead(node)
}

func (c *LRUCache) appendToHead(node *Node)  {
	if node == nil {
		return
	}

	node.Prev = nil
	node.Next = c.Head

	if node.Next != nil {
		node.Next.Prev = node
	}

	c.Head = node

	if c.Tail == nil {
		c.Tail = node
	}
}

func (c *LRUCache) removeNode(node *Node) {
	if node == nil {
		return
	}

	if c.Head == node {
		c.Head = node.Next
	}

	if c.Tail == node {
		c.Tail = node.Prev
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
}
