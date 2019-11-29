package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kalmeshbhavi/go-assignment/adapters/http"
	"github.com/kalmeshbhavi/go-assignment/engine"
	"github.com/kalmeshbhavi/go-assignment/providers/database"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS knights
(
    id SERIAL,
    name TEXT NOT NULL,
    strength int NOT NULL default 1,
    weapon_power int NOT NULL default 0,
    CONSTRAINT knights_pkey PRIMARY KEY (id)
)`

func EnsureTableExists(provider *database.Provider) {
	if _, err := provider.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func main() {

	//configs := config.GetConfigs()

	connString := database.GetConnectionString()
	provider := database.NewProvider(connString)
	EnsureTableExists(provider)

	e := engine.NewEngine(provider)

	adapter := http.NewHTTPAdapter(e)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	defer close(stop)

	adapter.Start()

	<-stop

	adapter.Stop()
	provider.Close()
}
