package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RaraykinValery/l0/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_USER"),
		"wildberries")
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
}

func ConnectToDB() (*sql.DB, error) {
	var err error

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_USER"),
		"wildberries")
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InsertOrderToDB(order models.Order) error {
	bOrder, err := json.Marshal(order)
	if err != nil {
		if err != nil {
			log.Printf("Failed to unmarshal order: %v", err)
			return err
		}
	}

	_, err = db.Exec(`INSERT INTO orders ("uuid", "data") values ($1, $2)`,
		order.OrderUID, bOrder)
	if err != nil {
		return err
	}

	return nil
}

func GetOrderFromDB(uuid string) (models.Order, error) {
	var bOrder []byte
	var order models.Order

	row := db.QueryRow("SELECT data FROM orders WHERE uuid = $1", uuid)
	err := row.Scan(&bOrder)
	if err != nil {
		return models.Order{}, err
	}

	err = json.Unmarshal(bOrder, &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil

}
