package cache

import (
	"log"

	"github.com/RaraykinValery/l0/internal/database"
	"github.com/RaraykinValery/l0/internal/models"
)

var app_cache = make(map[string]models.Order)

func Init() error {
	log.Printf("Loading orders to cache from database...")

	err := loadOrdersFromDBToCache()
	if err != nil {
		return err
	}

	log.Printf("%v orders have been loaded to cache.", len(app_cache))

	return nil
}

func GetOrder(uuid string) (models.Order, bool) {
	val, ok := app_cache[uuid]
	return val, ok
}

func PutOrder(order models.Order) {
	app_cache[order.OrderUID] = order
}

func loadOrdersFromDBToCache() error {
	orders, err := database.SelectAllOrders()
	if err != nil {
		return err
	}

	for _, v := range orders {
		PutOrder(v)
	}

	return nil
}
