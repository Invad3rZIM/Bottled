package bottles

import (
	bottle "bottled/bottles/bottle"
	bottles "bottled/bottles/bottle"
	"bottled/database"
	"bottled/points"
	"bottled/utils"
	"time"
)

type BottleCache struct {
	Bottles       map[string]map[string]map[int]*bottle.Bottle
	BottleSenders map[int]int
	BottlesMade   int
	DB            *database.DatabaseConnection
	AllBottles    map[int]*bottle.Bottle
	DBChanges     chan *bottles.Bottle
}

func (bm *BottleCache) Launch() {
	go bm.CleanDB()
}

func (bm *BottleCache) CleanDB() {
	for {
		bm.DB.DeleteLifelessBottles()

		//only sweep database once every 5 minutes
		time.Sleep(time.Duration(300) * time.Second)
	}
}

func UpdateDBChanges() {

}

func NewBottleCache(d *database.DatabaseConnection) *BottleCache {
	bm := BottleCache{
		Bottles:       make(map[string]map[string]map[int]*bottle.Bottle),
		BottleSenders: make(map[int]int),
		DB:            d,
		AllBottles:    make(map[int]*bottle.Bottle),
		DBChanges:     make(chan *bottle.Bottle, 100),
	}

	n1 := []string{"local", "global"}
	n2 := []string{"lolz", "#deep", "O.o", "stories"}

	for _, a := range n1 {
		bm.Bottles[a] = make(map[string]map[int]*bottle.Bottle)
		for _, b := range n2 {
			bm.Bottles[a][b] = make(map[int]*bottle.Bottle)
		}
	}

	//testing purposes
	bm.CreateLocalBottles()
	bm.CreateGlobalBottles()

	return &bm
}

func (bm *BottleCache) CreateBottle(senderID int, message string, tag string, lives int, point points.Point) *bottle.Bottle {
	b := bottle.NewBottle(senderID, message, tag, lives, bm.BottlesMade)
	point.BottleID = b.BottleID

	b.AddLocation(point)
	if point.Enabled {

		bm.Bottles["local"][tag][b.BottleID] = b
	} else {
		bm.Bottles["global"][tag][b.BottleID] = b
	}

	bm.BottleSenders[b.BottleID] = b.SenderID
	bm.BottlesMade = bm.BottlesMade + 1

	bm.DB.AddBottle(b)

	return b
}

//testing function
func (bm *BottleCache) CreateGlobalBottles() {

	tags := []string{"lolz", "#deep", "O.o", "stories"}

	for i := 0; i < 1000; i++ {
		bid := utils.GenInt(0, 999999)
		uid := utils.GenInt(0, 999999)

		b := bottle.Bottle{
			BottleID: bid,
			SenderID: uid,
			Tag:      tags[utils.GenInt(0, len(tags))],
			Point: points.Point{
				Age:      utils.GenInt(0, 99999),
				BottleID: bid,
				Enabled:  false,
			},
		}

		bm.Bottles["global"][b.Tag][b.BottleID] = &b
		bm.BottleSenders[bid] = uid
	}
}

//local testing function
func (bm *BottleCache) CreateLocalBottles() {

	tags := []string{"lolz", "#deep", "O.o", "stories"}

	for i := 0; i < 1000; i++ {
		bid := utils.GenInt(0, 999999)
		uid := utils.GenInt(0, 999999)

		b := bottle.Bottle{
			BottleID: bid,
			SenderID: uid,
			Tag:      tags[utils.GenInt(0, len(tags))],
			Point: points.Point{
				Age:      utils.GenInt(0, 99999),
				Distance: utils.GenFloat(0, 999),
				Lat:      utils.GenFloat(42, 48),
				Long:     utils.GenFloat(60, 63),
				BottleID: bid,
				Enabled:  true,
			},
		}

		bm.Bottles["local"][b.Tag][b.BottleID] = &b
		bm.BottleSenders[bid] = uid

		//		fmt.Printf("%v", bm.BottleSenders[bid])
	}
}

//need to implement ideal max distance better later
func (bm *BottleCache) GetLocalBottles(tag string, p points.Point, amount int, idealMaxDistance float64) []*bottle.Bottle {
	bs := []*bottle.Bottle{}

	for _, v := range bm.Bottles["local"][tag] {
		v.Point.Distance = p.CalcDistance(v.Point)
	}

	op := GetClosestPoints(bm.Bottles["local"][tag], amount, idealMaxDistance, tag)

	for i := 0; i < len(*op) && i < amount; i++ {
		bs = append(bs, bm.Bottles["local"][tag][(*op)[i].bottleID])
	}
	return bs
}

//gets all the oldest bottles in the global section
func (bm *BottleCache) GetGlobalBottles(tag string, p points.Point, amount int) []*bottle.Bottle {
	bs := []*bottle.Bottle{}

	op := GetOldestPoints(bm.Bottles["global"][tag], amount, tag)

	for i := 0; i < len(*op) && i < amount; i++ {
		bs = append(bs, bm.Bottles["global"][tag][(*op)[i].bottleID])
	}
	return bs
}
