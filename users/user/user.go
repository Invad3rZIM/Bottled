package user

import "errors"

type User struct {
	Name      string
	Caps      int //bottleCurrency
	MaxHearts int

	UserID int
	Pin    int

	Key FriendKey

	Modified bool
}

//make sure messages are legitimate via pin verification
func (u User) Verify(pin int) bool {
	return u.Pin == pin
}

//User constructor
func NewUser(name string, caps int, hearts int, uid int, fk int, pin int) *User {
	u := User{
		Name:      name,
		Caps:      caps,
		MaxHearts: hearts,
		UserID:    uid,
		Key:       FriendKey{UserID: uid, Key: fk, Name: name},
		Pin:       pin,
	}

	return &u
}

func (u *User) AddCap(a int) {
	u.Caps = u.Caps + a
}

func (u *User) SubCap(a int) error {
	if u.Caps < a {
		return errors.New("insufficient funds")
	}

	u.Caps = u.Caps - a

	return nil
}

func (u *User) Befriend(v *User) error {
	if err := u.SubCap(1); err != nil {
		return err
	}

	v.AddCap(2)

	return nil
}
