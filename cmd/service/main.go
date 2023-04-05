package main

import (
	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/database"
	"github.com/RaraykinValery/l0/internal/http_server"
	"github.com/RaraykinValery/l0/internal/subscriber"
)

func main() {
	var err error

	err = database.Connect()
	if err != nil {
		panic(err)
	}

	err = cache.Init()
	if err != nil {
		panic(err)
	}

	err = subscriber.StartSubscriber()
	if err != nil {
		panic(err)
	}

	err = http_server.StartHTTPServer(":8080")
	if err != nil {
		panic(err)
	}
}
