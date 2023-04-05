package config

import "os"

var Server = AppServer{
	PORT: ":8080",
}

var Database = AppDatabase{
	HOST:     "localhost",
	PORT:     "5432",
	USER:     os.Getenv("POSTGRES_USER"),
	PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
	DB_NAME:  "wildberries",
}
