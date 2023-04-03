package cache

import (
	"log"

	"github.com/RaraykinValery/l0/database"
	"github.com/RaraykinValery/l0/models"
)

var (
	app_cache        = Cache{Data: make(map[string]models.Order)}
	initialised bool = false
)

type Cache struct {
	Data map[string]models.Order
}

func init() {
	if initialised {
		return
	}

	err := LoadOrdersFromDBToCache()
	if err != nil {
		log.Printf("Couldn't load orders from database: %s", err.Error())
		panic(err)
	}
	log.Print("Orders have been loaded to cache.")
	log.Printf("Cache size = %v", len(app_cache.Data))

	initialised = true
}

func GetOrderFromCache(uuid string) (models.Order, bool) {
	val, ok := app_cache.Data[uuid]
	return val, ok
}

func PutOrderToCache(order models.Order) {
	app_cache.Data[order.OrderUID] = order
}

func LoadOrdersFromDBToCache() error {
	orders, err := database.GetAllOrdersFromDB()
	if err != nil {
		return err
	}

	for _, v := range orders {
		PutOrderToCache(v)
	}

	return nil
}
