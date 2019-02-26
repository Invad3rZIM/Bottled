package handlers

import (
	"math/rand"
	"time"
)

type User struct {
	firstName  string
	userID     int
	bottleCaps int //bottleCurrency

	*HeartContainer
}

func (u User) GetUserID() int {
	return u.userID
}

func (u User) GetName() string {
	return u.firstName
}

func (u User) CapCount() int {
	return u.bottleCaps
}

func NewUser(name string) *User {
	id := genUserID()
	u := User{
		firstName:  name,
		userID:     id,
		bottleCaps: 100,

		HeartContainer: NewHeartContainer(id, 3, 1),
	}

	return &u
}

//filler id generator function. needs rework later
func genUserID() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(999999999)
}
