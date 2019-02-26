package handlers

type HeapNode struct {
	Relations []*HeapNode //0 -> parent, 1 -> left, 2 -> right

	Distance     float64
	Age          int
	BottleID     int
	Instantiated int
}

type Heap struct {
	root      *HeapNode
	last      *HeapNode
	cap       int
	len       int
	heuristic int //heuristic = 0 if sorting by distance, 1 if sorting by age
}

func NewHeap(heuristic int, root *HeapNode, cap int) *Heap {
	h := Heap{
		root:      root,
		last:      root,
		cap:       cap,
		len:       1,
		heuristic: heuristic,
	}

	return &h
}

func (hn *Heap) AddNode(n *HeapNode) {
	last := hn.last

	if !last.HasLeft() {

	}
}

func (hn HeapNode) HasLeft() bool {
	return (hn.Left().Instantiated == 1)
}

func (hn HeapNode) HasRight() bool {
	return (hn.Right().Instantiated == 1)
}

func NewHeapNode(bottleID int, distance float64, age int) *HeapNode {
	hn := HeapNode{
		Relations:    make([]*HeapNode, 3, 3),
		Age:          age,
		Distance:     distance,
		BottleID:     bottleID,
		Instantiated: 1,
	}

	return &hn
}

func (hn *HeapNode) Left() *HeapNode {
	return hn.Relations[1]
}

func (hn *HeapNode) Right() *HeapNode {
	return hn.Relations[2]
}

func (h *Heap) DecreaseSize() {
	h.cap = h.cap - 1
}
