package handlers

import (
	"sync"
	"time"
)

//performs one full cycle of healing
func (h *Handler) HealAll(wg *sync.WaitGroup) {
	for len(h.BrokenHearts) > 0 {
		hc := <-h.BrokenHearts
		fullyHealed := hc.Heal()

		if !fullyHealed {
			h.NewlyHurt <- hc
		}
	}

	wg.Done()
}

func (h *Handler) AddWounded(userID int) {
	hc := h.UserCache[userID].HeartContainer
	h.NewlyHurt <- hc
}

func (h *Handler) MedicManager(workerCount int, regenTime int) {
	//infinite loop
	for {
		//add all the fresh wounds at the start of the cycle

		for len(h.NewlyHurt) > 0 {
			h.BrokenHearts <- <-h.NewlyHurt
		}

		var wg sync.WaitGroup

		//populate worker channels
		for i := 0; i < workerCount; i++ {
			wg.Add(1)
			go h.HealAll(&wg)
		}

		wg.Wait()

		time.Sleep(time.Duration(regenTime) * time.Second)
	}
}
