package database

import (
	hearts "bottled/hearts/heart"
	"fmt"
)

//adds new heart to the database
func (b *DatabaseConnection) AddHeart(h *hearts.Heart) error {
	sqlStatement := `
	INSERT INTO hearts (userid, current, max, rate)
	VALUES ($1, $2, $3,$4)`
	_, err := b.db.Exec(sqlStatement, h.UserID, h.Current, h.Max, h.Rate)

	if err != nil {
		return err
	}

	return nil
}

//updates heart value in database
func (b *DatabaseConnection) UpdateHearts(h *hearts.Heart) error {
	sqlStatement := `UPDATE hearts
	SET current = ($2) 
	WHERE userID = ($1);`
	_, err := b.db.Exec(sqlStatement, h.UserID, h.Current)

	if err != nil {
		return err
	}

	return nil
}

//gets all wounded hearts.. used for initializing stuff
func (b *DatabaseConnection) AllWounded() (*[]*hearts.Heart, error) {
	sqlStatement := fmt.Sprintf(`SELECT userid, current, max, rate  FROM hearts WHERE current < max`)

	hc := []*hearts.Heart{}

	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		h := hearts.Heart{}

		rows.Scan(&h.UserID, &h.Current, &h.Max, &h.Rate)

		hc = append(hc, &h)
	}

	return &hc, nil
}

//searches datbase for heart. if not in there, we make it...
func (b *DatabaseConnection) GetHeart(userid int) (*hearts.Heart, error) {
	sqlStatement := fmt.Sprintf(`SELECT userid, current, max, rate  FROM hearts WHERE userid=%d`, userid)
	h := hearts.Heart{}

	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&h.UserID, &h.Current, &h.Max, &h.Rate)
		if err != nil {
			panic(err)
		}
		return &h, nil
	}

	u, err := b.GetUser(userid)

	if err != nil {
		return nil, err
	}

	h = hearts.Heart{
		UserID:  u.UserID,
		Current: u.MaxHearts,
		Max:     u.MaxHearts,
		Rate:    1,
	}

	b.AddHeart(&h)

	return &h, nil
}
