package main

import (
	"github.com/RaraykinValery/l0/internal/http_server"
	"github.com/RaraykinValery/l0/internal/subscriber"
)

func main() {
	var err error

	sc, sub, err := subscriber.StartSubscriber()
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	defer sub.Unsubscribe()

	err = http_server.StartHTTPServer(":8080")
	if err != nil {
		panic(err)
	}
}
