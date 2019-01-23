type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
} 

func (h *IntHeap) Pop() interface{} {
    ret := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return ret
}

func findKthLargest(nums []int, k int) int {
    h := IntHeap{}
    heap.Init(&h)
    
    for _, v := range nums {
        heap.Push(&h, v)
        
        if h.Len() > k {
            heap.Pop(&h)
        }
    }
    
    return h[0]
}
