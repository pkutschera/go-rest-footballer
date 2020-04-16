package main

import (
	"database/sql"
	"errors"
)

type footballer struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Position  string `json:"position"`
	ClubId    int    `json:"clubid"`
}

func (p *footballer) getFootballer(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *footballer) updateFootballer(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *footballer) deleteFootballer(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *footballer) createFootballer(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getFootballers(db *sql.DB) ([]footballer, error) {
	rows, err := db.Query("SELECT id, firstname, lastname, position, clubid FROM footballer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := []footballer{}

	for rows.Next() {
		var player footballer
		if err := rows.Scan(&player.Id, &player.FirstName, &player.LastName, &player.Position, &player.ClubId); err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}
