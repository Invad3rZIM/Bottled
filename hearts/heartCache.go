package hearts

import (
	"bottled/database"
	h "bottled/hearts/heart"
	"errors"
	"fmt"
	"sync"
	"time"
)

type HeartCache struct {
	Hearts map[int]*h.Heart
	//	DB        *database.DatabaseConnection
	DBChanges chan *h.Heart

	//used for healing the hearts
	Wounded  chan *h.Heart
	Cooldown chan *h.Heart
}

func NewHeartCache(d *database.DatabaseConnection) *HeartCache {
	hc := HeartCache{
		Hearts: make(map[int]*h.Heart),
		//	DB:        d,
		DBChanges: make(chan *h.Heart, 100),

		Wounded:  make(chan *h.Heart, 100),
		Cooldown: make(chan *h.Heart, 100),
	}

	return &hc
}

func (hc *HeartCache) Launch() {
	go hc.HealFactory(5, 5)
	go hc.DatabaseUpdater()
	go hc.GetAllWoundedFromDatabase()
}

//infinite loop of database changes
func (hc *HeartCache) DatabaseUpdater() {
	for {
		//h := <-hc.DBChanges
		//	hc.DB.UpdateHearts(h)
		<-hc.DBChanges
	}
}

func (hc *HeartCache) GetAllWoundedFromDatabase() error {

	//	h, err := hc.DB.AllWounded()

	//if err != nil {
	//	return err
	//}

	//for _, v := range *h {
	//	hc.Wounded <- v
	//}

	//	return nil

	return errors.New("db not connected")
}

func (hc *HeartCache) Harm(userID int, amount int) error {
	h, err := hc.GetHeart(userID)

	if err != nil {
		return err
	}

	if h.Current < amount {
		return errors.New("insufficient hearts")
	}

	h.Current = h.Current - amount

	hc.Wounded <- h
	hc.DBChanges <- h

	go hc.AddWounded(h)

	return nil
}

func (hc *HeartCache) GetHeart(userID int) (*h.Heart, error) {

	if h, ok := hc.Hearts[userID]; ok {
		return h, nil
	}

	//grabs heart from database
	//h, err := hc.DB.GetHeart(userID)
	err := errors.New("database not connected")
	//if heart not in database
	if err != nil {
		return nil, err
	}

	//hc.Hearts[userID] = h

	//if heart needs healing
	//if h.Current < h.Max {
	//	hc.AddWounded(h)
	//}

	//	return h, nil

	return nil, err
}

func (hc *HeartCache) AddWounded(h *h.Heart) {
	hc.Wounded <- h
}

//advances all wounded hearts by 1 healing cycle
func (hc *HeartCache) HealAll(wg *sync.WaitGroup) {
	fmt.Println(len(hc.Wounded))
	for len(hc.Wounded) > 0 {
		heart := <-hc.Wounded

		if heart.OnCooldown == false {

			fullyHealed := heart.Heal(1)

			if !fullyHealed {
				heart.OnCooldown = true
				hc.Cooldown <- heart

			}

			//update database channel
			hc.DBChanges <- heart
		}

	}

	wg.Done()
}

func (hc *HeartCache) HealFactory(workerCount int, regenTime int) {
	//infinite loop
	for {

		fmt.Println("XXX")

		//add all the fresh wounds at the start of the cycle
		for len(hc.Cooldown) > 0 {
			heart := <-hc.Cooldown
			heart.OnCooldown = false
			hc.Wounded <- heart
		}

		var wg sync.WaitGroup

		//populate worker channels
		for i := 0; i < workerCount; i++ {
			wg.Add(1)
			go hc.HealAll(&wg)
		}

		wg.Wait()

		//cooldown for next time
		time.Sleep(time.Duration(regenTime) * time.Second)

	}
}
