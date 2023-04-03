package subscriber

import (
	"encoding/json"
	"log"

	"github.com/RaraykinValery/l0/cache"
	"github.com/RaraykinValery/l0/database"
	"github.com/RaraykinValery/l0/models"
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

func StartSubscriber() {
	sc, err := stan.Connect("test-cluster", "client-subscriber-1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		panic(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("orders",
		messageHandler,
		stan.DurableName("my-durable"))
	if err != nil {
		log.Fatalf("Failed to subscribe to subject: %v", err)
		panic(err)
	}
	defer sub.Close()

	select {}
}
