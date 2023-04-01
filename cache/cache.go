package cache

import (
	"errors"

	"github.com/RaraykinValery/l0/models"
)

var app_cache Cache

type Cache struct {
	Data map[string]models.Order
}

func GetOrderFromCache(uuid string) (models.Order, bool) {
	val, ok := app_cache.Data[uuid]
	return val, ok
}

func PutOrderToCache(order models.Order) {
	app_cache.Data[order.OrderUID] = order
}

func LoadOrdersFromDBToCache() error {
	return errors.New("error")
}

func init() {
	app_cache = Cache{
		Data: make(map[string]models.Order),
	}
}
