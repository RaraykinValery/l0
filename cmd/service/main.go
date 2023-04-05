package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/RaraykinValery/l0/internal/service"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := service.Start()
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	err = service.Stop()
	if err != nil {
		panic(err)
	}
}
