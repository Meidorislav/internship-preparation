type Item struct {
	val  int
	freq int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
    counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for val, freq := range counts {
		heap.Push(h, Item{val: val, freq: freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		res = append(res, item.val)
	}

	return res
}

