package subscriber

import (
	"encoding/json"
	"log"

	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/database"
	"github.com/RaraykinValery/l0/internal/models"
	"github.com/nats-io/stan.go"
)

var sc stan.Conn
var sub stan.Subscription

func messageHandler(msg *stan.Msg) {
	var order models.Order

	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		log.Printf("Failed to unmarshal order: %s", err)
		return
	}

	cache.PutOrder(order)

	err = database.InsertOrder(order)
	if err != nil {
		log.Printf("Failed to write order to database: %s", err)
		return
	}

	log.Printf("Received order with uid: %v", order.OrderUID)
}

func Start() error {
	var err error

	sc, err = stan.Connect(
		"test-cluster",
		"client-subscriber-1",
		stan.NatsURL("nats://nats-streaming:4222"),
	)
	if err != nil {
		return err
	}

	sub, err = sc.Subscribe(
		"orders",
		messageHandler,
	)
	if err != nil {
		return err
	}

	return nil
}

func Stop() error {
	err := sub.Unsubscribe()
	if err != nil {
		return err
	}

	err = sc.Close()
	if err != nil {
		return err
	}

	return nil
}
