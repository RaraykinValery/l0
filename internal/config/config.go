package config

var Server = AppServer{
	PORT: ":8080",
}

var Database = AppDatabase{
	HOST:     "database",
	PORT:     "5432",
	USER:     "dbuser",
	PASSWORD: "pass",
	DB_NAME:  "wb",
	SSLMODE:  "disable",
}
