package handlers

import (
	"bottled/bottles"
	"bottled/database"

	"bottled/hearts"
	"bottled/messages"
	"bottled/users"
)

type Handler struct {
	//	NewlyHurt     chan *HeartContainer
	//	BrokenHearts  chan *HeartContainer
	UserCache    *users.UserCache
	BottleCache  *bottles.BottleCache
	MessageCache *messages.MessageCache
	HeartCache   *hearts.HeartCache

	DB *database.DatabaseConnection
}

func NewHandler(d *database.DatabaseConnection) *Handler {
	var h Handler

	h.UserCache = users.LoadUserCache(d) //needs sql hookup
	h.MessageCache = messages.NewMessageCache(d)
	h.BottleCache = bottles.NewBottleCache(d)
	h.HeartCache = hearts.NewHeartCache(d)

	h.DB = d

	return &h
}

/*
func (h *Handler) AddUserCache(u *User) {
	h.UserCache[u.GetUserID()] = u
}

func (h *Handler) CreateUserToDatabase(name string) *User {

	u := NewUser(name)

	h.AddUserCache(u)

	return u
}

func (h *Handler) Hurt(userID int, pain int) {
	u := h.UserCache[userID]
	u.Hurt(pain)
	h.AddWounded(u.GetUserID())
}
*/
