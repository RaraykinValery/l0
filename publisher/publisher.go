package main

import (
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster",
		"client-publisher-1",
		stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer sc.Close()

	for {
		time.Sleep(5 * time.Second)
		sc.Publish("orders", []byte("Hello World"))
	}
}
