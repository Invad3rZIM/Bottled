package heart

import (
	u "bottled/users/user"
)

type Heart struct {
	UserID     int
	Current    int
	Max        int
	Rate       int //# of fragments you regenerate each cycle
	OnCooldown bool
}

func NewHeart(u *u.User) *Heart {
	frags := u.MaxHearts * 4

	hc := Heart{
		UserID:  u.UserID,
		Current: frags,
		Max:     frags,
		Rate:    1,
	}

	return &hc
}

//heals heart by specified amount. returns true if fully healed
func (h *Heart) Heal(amount int) bool {

	//only heals if not already healed that cycle
	if !h.OnCooldown {
		h.Current = h.Current + amount

		if h.Current < h.Max {
			return false
		}

		if h.Current > h.Max {
			h.Current = h.Max
			return true
		}

		return true

	}

	return false
}
