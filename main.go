package main

import (
	"github.com/kalmeshbhavi/go-assignment/adapters/http"
	"github.com/kalmeshbhavi/go-assignment/engine"
	"github.com/kalmeshbhavi/go-assignment/providers/database"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	provider := database.NewProvider()
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
