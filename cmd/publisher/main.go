package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/RaraykinValery/l0/internal/models"
	stan "github.com/nats-io/stan.go"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func readJSONSample() (models.Order, error) {
	var orderData []byte
	var order models.Order

	orderData, err := os.ReadFile("model.json")
	if err != nil {
		return models.Order{}, err
	}

	err = json.Unmarshal(orderData, &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func RandStringRunes(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func main() {
	order, err := readJSONSample()
	if err != nil {
		log.Fatal("Couldn't read data from model.json")
	}

	sc, err := stan.Connect("test-cluster",
		"client-publisher-1",
		stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer sc.Close()

	rand.Seed(time.Now().UnixNano())

	for {
		time.Sleep(10 * time.Second)
		order.OrderUID = RandStringRunes(15) + "test"
		bOrder, err := json.Marshal(order)
		if err != nil {
			log.Fatalf("Failed to marshal order: %v", err)
		}

		sc.Publish("orders", bOrder)
	}
}
