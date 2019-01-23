// Copy from golang doc
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type KthLargest struct {
	Heap *IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	intHeap := IntHeap{}
	heap.Init(&intHeap)

	// push all elements to the heap
	for _, v := range nums {
		heap.Push(&intHeap, v)
	}

	// remove the smaller elements from the heap such that only the k largest elements are in the heap
	for len(intHeap) > k {
		heap.Pop(&intHeap)
	}

	return KthLargest{Heap: &intHeap, K: k}
}


func (c *KthLargest) Add(val int) int {
	heap.Push(c.Heap, val)

	if c.Heap.Len() > c.K {
		heap.Pop(c.Heap)
	}

	return (*c.Heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
