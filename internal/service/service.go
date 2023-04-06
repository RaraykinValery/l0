package service

import (
	"log"

	"github.com/RaraykinValery/l0/internal/cache"
	"github.com/RaraykinValery/l0/internal/config"
	"github.com/RaraykinValery/l0/internal/database"
	"github.com/RaraykinValery/l0/internal/http_server"
	"github.com/RaraykinValery/l0/internal/subscriber"
)

func Start() error {
	err := database.Connect()
	if err != nil {
		return err
	}

	err = database.CreateTables()
	if err != nil {
		return err
	}

	err = cache.Init()
	if err != nil {
		return err
	}

	err = subscriber.Start()
	if err != nil {
		return err
	}

	err = http_server.Start(config.Server.PORT)
	if err != nil {
		return err
	}

	return nil
}

func Stop() error {
	log.Print("Shutting down the service...")

	err := database.Disconnect()
	if err != nil {
		return err
	}

	err = subscriber.Stop()
	if err != nil {
		return err
	}

	err = http_server.Shutdown()
	if err != nil {
		return err
	}

	log.Print("Service is down")

	return nil
}
