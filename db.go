package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func Connect(c *Config) (*pgx.Conn, error) {

	// Create the url to the db
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.host, c.port, c.user, c.password, c.dbname)

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
