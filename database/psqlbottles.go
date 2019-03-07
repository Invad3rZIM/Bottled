package database

import (
	bottles "bottled/bottles/bottle"
	"fmt"
	"math/rand"
	"time"
)

//adds the bottle to the database, generating a unique bottle id along the way
func (d *DatabaseConnection) AddBottle(b *bottles.Bottle) error {
	b.BottleID = d.GenUniqueBottleID()

	sqlStatement := `
	INSERT INTO BOTTLES (bottleid, senderid, message, tag, lives, age, enabled, lat, long)
	VALUES ($1, $2, $3,$4, $5, $6, $7, $8, $9)`
	_, err := d.db.Exec(sqlStatement, b.BottleID, b.SenderID, b.Message, b.Tag, b.Lives, b.Age, b.Point.Enabled, b.Point.Lat, b.Point.Long)

	if err != nil {
		return err
	}

	return nil
}

func (b *DatabaseConnection) GetBottleCount() int {
	var cnt int
	b.db.QueryRow(`select count(*) from BOTTLES; `).Scan(&cnt)

	return cnt
}

func (b *DatabaseConnection) GenUniqueBottleID() int {
	var bid int

	for {
		bid = GenBottleID()
		sqlStatement := fmt.Sprintf(`SELECT bottleID FROM users WHERE userid=%d`, bid)

		// Replace 3 with an ID from your database or another random
		// value to test the no rows use case.
		rows, _ := b.db.Query(sqlStatement)
		flag := true

		if rows == nil {
			return bid
		}

		for rows.Next() {
			rows.Close()
			flag = false
		}

		if flag == true {
			break
		}
	}

	return bid
}

//remap to make sure bottlesID is unique
func GenBottleID() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(999999999)
}

//updates caps value in user database
func (d *DatabaseConnection) UpdateBottleLives(b *bottles.Bottle) error {
	sqlStatement := `UPDATE bottles
	SET lives = ($2) 
	WHERE bottleID = ($1);`
	_, err := d.db.Exec(sqlStatement, b.BottleID, b.Lives)

	if err != nil {
		return err
	}

	return nil
}

//kills all entrys with zero lives left
func (d *DatabaseConnection) DeleteLifelessBottles() error {
	sqlStatement := `DELETE from bottles where lives = 0;`

	_, err := d.db.Exec(sqlStatement)

	if err != nil {
		return err
	}

	return nil
}
