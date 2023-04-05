package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/RaraykinValery/l0/internal/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		"wildberries")
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	log.Print("Connected to database")

	return nil
}

func InsertOrder(order models.Order) error {
	bOrder, err := json.Marshal(order)
	if err != nil {
		log.Printf("Failed to unmarshal order: %v", err)
		return err
	}

	_, err = db.Exec(`INSERT INTO orders ("uuid", "data") values ($1, $2)`,
		order.OrderUID, bOrder)
	if err != nil {
		return err
	}

	return nil
}

func SelectOrder(uuid string) (models.Order, error) {
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

func SelectAllOrders() ([]models.Order, error) {
	var orders []models.Order

	rows, err := db.Query("SELECT data FROM orders")
	if err != nil {
		return []models.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var data []byte
		var order models.Order

		if err := rows.Scan(&data); err != nil {
			return []models.Order{}, err
		}

		err := json.Unmarshal(data, &order)
		if err != nil {
			return []models.Order{}, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
