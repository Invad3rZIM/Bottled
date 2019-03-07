package users

import (
	"bottled/database"
	u "bottled/users/user"
	"bottled/utils"
	"errors"
	"fmt"
)

type UserCache struct {
	Users        map[int]*u.User
	DB           *database.DatabaseConnection
	FriendCodes  map[int][]*u.FriendKey
	DBCapChanges chan *u.User
}

func NewUserCache(d *database.DatabaseConnection) *UserCache {
	uc := UserCache{
		Users:        make(map[int]*u.User),
		DB:           d,
		FriendCodes:  make(map[int][]*u.FriendKey),
		DBCapChanges: make(chan *u.User),
	}

	go uc.DatabaseUpdater()

	return &uc
}

func (uc UserCache) GetUser(userID int) (*u.User, error) {
	if user, ok := uc.Users[userID]; ok {
		return user, nil
	}

	//pull from database
	user, err := uc.DB.GetUser(userID)

	if err != nil {
		return nil, err
	}

	uc.Users[user.UserID] = user

	//instantiate the friend mailbox
	if _, ok := uc.FriendCodes[userID]; !ok {
		uc.FriendCodes[userID] = []*u.FriendKey{}
	}

	return user, nil

}

//infinite loop of database changes
func (uc *UserCache) DatabaseUpdater() {
	for {
		uc.DB.UpdateCaps(<-uc.DBCapChanges)
	}
}

//wrapper to add caps to a user
func (uc *UserCache) AddCaps(userID int, quantity int) error {
	u, err := uc.GetUser(userID)

	if err != nil {
		return err
	}

	u.AddCap(quantity)

	uc.DBCapChanges <- u

	fmt.Println("DDDDDDD")

	return nil
}

//gives out friendkeys to the friendkey mailbox of the recipient
func (uc *UserCache) GiveFriendKey(sid int, rid int) error {
	s, err := uc.GetUser(sid)

	if err != nil {
		return err
	}

	r, err := uc.GetUser(rid)

	if err != nil {
		return err
	}

	//try to give 1 cap to increment theirs by 2
	err = s.Befriend(r)

	if err != nil {
		return err
	}

	uc.DBCapChanges <- s
	uc.DBCapChanges <- r

	return nil
}

//gets all the new friendkeys in your mailbox
func (uc *UserCache) GetFriendKeys(sid int) (*[]*u.FriendKey, error) {
	f, ok := uc.FriendCodes[sid]

	if !ok {
		return nil, errors.New("user not found")
	}

	return &f, nil
}

func LoadUserCache(d *database.DatabaseConnection) *UserCache {
	return NewUserCache(d)
}

func (uc UserCache) ValidPin(userID int, pin int) bool {
	u, err := uc.GetUser(userID)

	fmt.Printf("xxx%v", u)

	if err != nil {
		return false
	}

	return u.Pin == pin
}

//Testing Cache
func NewUserTestCache() *UserCache {
	uc := UserCache{
		Users: make(map[int]*u.User),
	}

	//testing purposes
	uc.CreateTestUsers()

	return &uc
}

//creates 1000 dummy users
func (uc *UserCache) CreateTestUsers() {
	for i := 0; i < 1000; i++ {
		user := u.User{
			UserID: utils.GenInt(0, 99999999),
			Pin:    utils.GenInt(0, 99999999),
		}

		user.Key = u.FriendKey{UserID: user.UserID, Key: utils.GenInt(0, 99999999)}

		uc.Users[user.UserID] = &user
	}
}
