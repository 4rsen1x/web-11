package provider

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

func (p *Provider) AddUser(username, password string) error {
	_, err := p.conn.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) GetUser(username string) (string, error) {
	var password string
	err := p.conn.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}
	return password, nil
}
