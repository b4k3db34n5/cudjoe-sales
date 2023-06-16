package main

import "context"

func main() {
	// Init Config
	c := Init()

	// Connect to DB
	conn, err := Connect(c)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	// Run GraphQL
	err = Run()
	if err != nil {
		panic(err)
	}
}
