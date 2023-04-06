package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/RaraykinValery/l0/internal/config"
	"github.com/RaraykinValery/l0/internal/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error

	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.HOST,
		config.Database.PORT,
		config.Database.USER,
		config.Database.PASSWORD,
		config.Database.DB_NAME,
		config.Database.SSLMODE,
	)

	log.Print("Connecting to database...")

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Print("Database connected")

	return nil
}

func Disconnect() error {
	err := db.Close()
	if err != nil {
		return err
	}

	return nil
}

func InsertOrder(order models.Order) error {
	bOrder, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO orders ("uuid", "data") values ($1, $2)`,
		order.OrderUID, bOrder,
	)
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
	var data []byte
	var order models.Order
	var all_orders []models.Order

	rows, err := db.Query("SELECT data FROM orders")
	if err != nil {
		return []models.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&data); err != nil {
			return []models.Order{}, err
		}

		err := json.Unmarshal(data, &order)
		if err != nil {
			return []models.Order{}, err
		}
		all_orders = append(all_orders, order)
	}

	return all_orders, nil
}

func CreateTables() error {
	log.Print("Creating database tables...")

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders (uuid VARCHAR(20) PRIMARY KEY, data JSONB NOT NULL)")
	if err != nil {
		return err
	}

	log.Print("Database tables have been created")

	return nil
}
