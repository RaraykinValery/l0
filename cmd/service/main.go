package main

import (
	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/database"
	"github.com/RaraykinValery/l0/internal/http_server"
	"github.com/RaraykinValery/l0/internal/subscriber"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var err error

	err = database.Connect()
	panicOnError(err)

	err = cache.Init()
	panicOnError(err)

	err = subscriber.Start()
	panicOnError(err)

	err = http_server.Start(":8080")
	panicOnError(err)
}
