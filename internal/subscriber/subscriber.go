package subscriber

import (
	"encoding/json"
	"log"

	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/database"
	"github.com/RaraykinValery/l0/internal/models"
	"github.com/nats-io/stan.go"
)

func messageHandler(msg *stan.Msg) {
	var order models.Order

	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		log.Printf("Failed to unmarshal order: %s", err)
		return
	}

	cache.PutOrderToCache(order)

	err = database.InsertOrderToDB(order)
	if err != nil {
		log.Printf("Failed to write order to database: %s", err)
		return
	}

	log.Printf("Received order with uid: %v", order.OrderUID)
}

func StartSubscriber() (stan.Conn, stan.Subscription, error) {
	sc, err := stan.Connect("test-cluster", "client-subscriber-1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		return nil, nil, err
	}

	sub, err := sc.Subscribe("orders",
		messageHandler,
		stan.DurableName("my-durable"))
	if err != nil {
		log.Fatalf("Failed to subscribe to subject: %v", err)
		return nil, nil, err
	}

	return sc, sub, nil
}
