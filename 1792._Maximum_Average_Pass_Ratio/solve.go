package main

import "container/heap"

type Item struct {
	Pass, Total int
	AverageDiff float64
}

type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].AverageDiff > h[j].AverageDiff
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	new_heap := make(Heap, 0, len(classes))
	for _, v := range classes {
		pass, total := v[0], v[1]
		average := float64(pass) / float64(total)
		new_average := float64(pass+1) / float64(total+1)
		item := Item{
			Total:       total,
			Pass:        pass,
			AverageDiff: new_average - average,
		}
		new_heap.Push(&item)
	}

	heap.Init(&new_heap)

	for extraStudents > 0 {
		min_item := heap.Pop(&new_heap).(*Item)
		min_item.Total += 1
		min_item.Pass += 1
		min_item.AverageDiff = float64(min_item.Pass+1)/float64(min_item.Total+1) - float64(min_item.Pass)/float64(min_item.Total)
		heap.Push(&new_heap, min_item)

		extraStudents -= 1
	}

	average_sum := 0.0
	for _, v := range new_heap {
		average_sum += float64(v.Pass) / float64(v.Total)
	}

	return average_sum / float64(len(new_heap))
}
