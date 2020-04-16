package main

import (
	"database/sql"
	"errors"
)

type club struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Founded int    `json:"founded"`
}

func (p *club) getClub(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *club) updateClub(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *club) deleteClub(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *club) createClub(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getClubs(db *sql.DB, start, count int) ([]club, error) {
	return nil, errors.New("Not implemented")
}
