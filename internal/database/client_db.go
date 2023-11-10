package database

import (
	"database/sql"

	"github.com/williamrlbrito/walletcore/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{DB: db}
}

func (db *ClientDB) Get(id string) (*entity.Client, error) {
	client := entity.Client{}
	smtm, err := db.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer smtm.Close()
	row := smtm.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}
	return &client, nil
}

func (db *ClientDB) Save(client entity.Client) error {
	smtm, err := db.DB.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer smtm.Close()
	_, err = smtm.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
