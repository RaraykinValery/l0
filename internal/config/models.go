package config

type AppServer struct {
	PORT string
}

type AppDatabase struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	DB_NAME  string
}
