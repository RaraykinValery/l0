package main

import (
	"log"

	"github.com/nats-io/stan.go"
)

func messageHandler(msg *stan.Msg) {
	log.Printf("Received message: %s", string(msg.Data))
}

func main() {
	sc, err := stan.Connect("test-cluster", "client-subscriber-1", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("orders", messageHandler, stan.DurableName("my-durable"))
	if err != nil {
		log.Fatalf("Failed to subscribe to subject: %v", err)
	}
	defer sub.Close()

	select {}
}
