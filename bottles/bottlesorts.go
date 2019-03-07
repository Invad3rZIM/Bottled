package bottles

import (
	bottle "bottled/bottles/bottle"
	"container/heap"
)

func GetClosestPoints(ps map[int]*bottle.Bottle, needed int, limit float64, bottleType string) *[]Node {
	cs := []Node{}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(ps))

	i := 0
	for _, p := range ps {
		pq[i] = &Node{
			bottleID: p.Point.BottleID,
			priority: int(p.Point.Distance),
			tag:      p.Tag,
		}
		i++
	}

	heap.Init(&pq) //empty pq

	count := 0

	// Take the items out; they arrive in decreasing priority order.
	for a := 0; a < len(pq) && count < needed; a++ {
		item := heap.Pop(&pq).(*Node)

		if item.tag == bottleType {
			count++
			cs = append(cs, *item)
		}

	}

	return &cs
}

func GetOldestPoints(ps map[int]*bottle.Bottle, needed int, bottleType string) *[]Node {
	cs := []Node{}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(ps))

	i := 0
	for _, p := range ps {
		pq[i] = &Node{
			bottleID: p.Point.BottleID,
			priority: int(p.Point.Age),
			tag:      p.Tag,
		}
		i++
	}

	heap.Init(&pq) //empty pq

	count := 0

	// Take the items out; they arrive in decreasing priority order.
	for a := 0; a < len(pq) && count < needed; a++ {
		item := heap.Pop(&pq).(*Node)

		if item.tag == bottleType {
			count++
			cs = append(cs, *item)
		}

	}

	return &cs
}

// An Item is something we manage in a priority queue.
type Node struct {
	priority int // The priority of the item in the queue.

	// The index is needed by update and is maintained by the heap.Interface methods.
	userID   int // The index of the item in the heap.
	bottleID int

	tag string

	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.

	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
