package handlers

type BottleManager struct {
	localBottles  map[int]*Bottle
	globalBottles map[int]*Bottle
	bottlesMade   int
}

func NewBottleManager() *BottleManager {
	bm := BottleManager{
		localBottles:  make(map[int]*Bottle),
		globalBottles: make(map[int]*Bottle),
	}

	return &bm
}

/*
func (bm *BottleManager) AcceptBottle(id int) int {

	b, ok := bm.localBottles[id]

	if ok {
		b.LoseLife(1) //accepting a bottle costs 1 life
		return b.SenderID
	}

	b, ok = bm.globalBottles[id]

	if ok {
		b.LoseLife(1)
		return b.SenderID
	}

	return b.SenderID
}
*/

type StatusMessage struct {
	Status int
}
