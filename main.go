package main

import (
	"github.com/RaraykinValery/l0/http_server"
	"github.com/RaraykinValery/l0/subscriber"
)

func main() {
	go subscriber.StartSubscriber()
	http_server.StartHTTPServer(":8080")
}
