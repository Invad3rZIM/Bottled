package handlers

import "net/http"

type Handler struct {
	UserCache     map[int]*User
	NewlyHurt     chan *HeartContainer
	BrokenHearts  chan *HeartContainer
	bottleManager *BottleManager
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	h.CreateUser("TESTING")
}

func NewHandler() *Handler {
	var h Handler

	h.UserCache = make(map[int]*User)
	h.NewlyHurt = make(chan *HeartContainer, 100000)
	h.BrokenHearts = make(chan *HeartContainer, 100000)
	h.bottleManager = NewBottleManager()

	return &h
}

func (h *Handler) AddUserCache(u *User) {
	h.UserCache[u.GetUserID()] = u
}

func (h *Handler) CreateUser(name string) *User {

	u := NewUser(name)

	h.AddUserCache(u)

	return u
}

func (h *Handler) Hurt(userID int, pain int) {
	u := h.UserCache[userID]
	u.Hurt(pain)
	h.AddWounded(u.GetUserID())
}
