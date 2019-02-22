package handlers

import (
	"math/rand"
	"time"
)

/*

	ai
	how technology changes our lives
	kermit the frog
	the meaning of intelligence
	what is real
	what is life

	i'm running out of words

	inspriation comes in waves...

	i'd have to think about it, play around with things a little bit.

	i'm not sure yet, but i'm going to keep writing the things that i'm thinking to imrpove my stream of consciousness abilities. i'm also typing things.

	simulacrum

	ethics of ai

	meerrrrrr
*/

type Bottle struct {
	bottleID int
	senderID int
	message  string
	lives    int
}

func NewBottle(senderID int, message string, lives int) *Bottle {
	b := Bottle{
		bottleID: GenBottleID(),
		senderID: senderID,
		message:  message,
		lives:    lives,
	}

	return &b
}

//remap to make sure bottlesID is unique
func GenBottleID() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(999999999)
}
