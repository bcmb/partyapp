package party

import (
	"context"
	"gotraining.com/database"
	"log"
	"time"
)

func InsertParty(party *Party) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO party (name, city, address, date_time) VALUES ($1,$2,$3,$4)`
	results, err := database.DbConn.ExecContext(
		ctx,
		query,
		party.Name,
		party.City,
		party.Address,
		party.DateTime)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if ra, _ := results.RowsAffected(); ra != 1 {
		log.Println(err.Error())
		return err
	}

	return nil
}

func GetParty(partyId int) (*Party, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT  id, name, city, address, date_time FROM party WHERE id = $1`
	result := database.DbConn.QueryRowContext(ctx, query, partyId)

	if err := result.Err(); err != nil {
		log.Println(err)
	}

	party := Party{}

	result.Scan(
		&party.Id,
		&party.Name,
		&party.City,
		&party.Address,
		&party.DateTime)

	return &party, nil
}

func GetParties() (*[]Party, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT  id, name, city, address, date_time FROM party WHERE date_time BETWEEN now() AND now() + '10y'`
	results, err := database.DbConn.QueryContext(ctx, query)

	if err != nil {
		log.Println(err.Error())
	}

	parties := make([]Party, 0)

	for results.Next() {
		var party Party
		results.Scan(
			&party.Id,
			&party.Name,
			&party.City,
			&party.Address,
			&party.DateTime)

		parties = append(parties, party)
	}

	return &parties, nil
}

func UpdateParty(party *Party, partyId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `UPDATE party SET "name" = $1, city = $2, address = $3, date_time = $4 WHERE id = $5;`
	_, err := database.DbConn.ExecContext(
		ctx,
		query,
		party.Name,
		party.City,
		party.Address,
		party.DateTime,
		partyId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func DeleteParty(partyId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM party WHERE id = $1;`
	_, err := database.DbConn.ExecContext(
		ctx,
		query,
		partyId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
