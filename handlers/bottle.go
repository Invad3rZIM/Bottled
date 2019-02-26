package handlers

import (
	"math/rand"
	"time"
)

type Bottle struct {
	bottleID   int
	SenderID   int    `json:"senderID"`
	Message    string `json:"message"`
	BottleType string
	Lives      int `json:"lives"`
	Point
}

func (b *Bottle) AddLocation(p Point) {
	b.Point = p
}

func NewBottle(senderID int, message string, bottleType string, lives int) *Bottle {
	b := Bottle{
		bottleID:   GenBottleID(),
		SenderID:   senderID,
		Message:    message,
		Lives:      lives,
		BottleType: bottleType,
	}

	return &b
}

func (b Bottle) GetBottleID() int {
	return b.bottleID
}

func (b *Bottle) LoseLife(takeaway int) bool {
	b.Lives = b.Lives - takeaway

	return b.Lives > 0
}

//remap to make sure bottlesID is unique
func GenBottleID() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(999999999)
}
