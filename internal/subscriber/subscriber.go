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
	sc, err := stan.Connect("test-cluster", "client-subscriber-1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		return err
	}

	sub, err = sc.Subscribe("orders",
		messageHandler,
		stan.DurableName("my-durable"))
	if err != nil {
		log.Fatalf("Failed to subscribe to subject: %v", err)
		return err
	}

	return nil
}
