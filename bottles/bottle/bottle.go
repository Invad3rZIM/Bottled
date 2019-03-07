package bottles

import (
	"bottled/points"
)

type Bottle struct {
	BottleID int
	SenderID int    `json:"senderID"`
	Message  string `json:"message"`
	Tag      string
	Lives    int `json:"lives"`
	Age      int
	Modified int //0 = no change, 1 = lives change

	points.Point
}

func (b *Bottle) AddLocation(p points.Point) {
	b.Point = p
}

func NewBottle(senderID int, message string, tag string, lives int, age int) *Bottle {
	b := Bottle{
		SenderID: senderID,
		Message:  message,
		Lives:    lives,
		Tag:      tag,
		Age:      age,
		Modified: 0,
	}

	return &b
}

func (b *Bottle) LoseLife() bool {
	b.Lives = b.Lives - 1
	b.Modified = 1

	return b.Lives > 0
}
