package main

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func Init() *Config {
	c := Config{
		host:     "localhost",
		port:     "5432",
		user:     "beans",
		password: "password",
		dbname:   "cs_db",
	}

	return &c
}
