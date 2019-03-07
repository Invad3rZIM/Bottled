package database

import (
	users "bottled/users/user"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func (b *DatabaseConnection) AddUser(u *users.User) error {
	u.UserID = b.GetUserCount() + 1
	u.Pin = genPin()
	u.Key = users.FriendKey{Key: genPin(), UserID: u.UserID, Name: u.Name}
	sqlStatement := `
	INSERT INTO users (userid, pin, friendkey, name, caps, maxHearts)
	VALUES ($1, $2, $3,$4, $5, $6)`
	_, err := b.db.Exec(sqlStatement, u.UserID, u.Pin, u.Key.Key, u.Name, u.Caps, u.MaxHearts)

	if err != nil {
		return err
	}

	return nil
}

//updates caps value in user database
func (b *DatabaseConnection) UpdateCaps(u *users.User) error {
	sqlStatement := `UPDATE users
	SET caps = ($2) 
	WHERE userID = ($1);`
	_, err := b.db.Exec(sqlStatement, u.UserID, u.Caps)

	if err != nil {
		return err
	}

	return nil
}

func (b *DatabaseConnection) GetUser(userid int) (*users.User, error) {
	sqlStatement := fmt.Sprintf(`SELECT userid, name, pin, friendkey, caps, maxhearts  FROM users WHERE userid=%d`, userid)
	u := users.User{}

	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&u.UserID, &u.Name, &u.Pin, &u.Key.Key, &u.Caps, &u.MaxHearts)
		if err != nil {
			panic(err)
		}
		return &u, nil
	}
	return nil, errors.New("error: user not found")
}

func (b *DatabaseConnection) GetUserCount() int {
	var cnt int
	b.db.QueryRow(`select count(*) from USERS; `).Scan(&cnt)

	return cnt
}

//filler id generator function. needs rework later
func genPin() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(99999999)
}
