package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/RaraykinValery/l0/models"
	stan "github.com/nats-io/stan.go"
)

var randomOrder models.Order = models.Order{
	OrderUID:    "b563feb7b2b84b6test",
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: models.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: models.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	},
	Items: []models.Item{
		{
			ChrtID:      9934930,
			TrackNumber: "WBILMTESTTRACK",
			Price:       453,
			RID:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NMID:        2389212,
			Brand:       "Vivienne Sabo",
			Status:      202,
		},
	},
	Locale:            "en",
	InternalSignature: "",
	CustomerID:        "test",
	DeliveryService:   "meest",
	ShardKey:          "9",
	SMID:              99,
	DateCreated:       time.Date(2021, 11, 26, 6, 22, 19, 0, time.UTC),
	OOFShard:          "1",
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	sc, err := stan.Connect("test-cluster",
		"client-publisher-1",
		stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		panic(err)
	}
	defer sc.Close()

	for {
		time.Sleep(10 * time.Second)
		randomOrder.OrderUID = RandStringRunes(15) + "test"
		bOrder, err := json.Marshal(randomOrder)
		if err != nil {
			log.Printf("Failed to marshal order: %v", err)
			return
		}

		sc.Publish("orders", bOrder)
	}
}
