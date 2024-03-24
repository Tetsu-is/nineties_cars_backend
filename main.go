package main

import (
	"context"
	"net/http"
	"nineties/db"
	"nineties/handler/router"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	const (
		defaultPort   = "8080"
		defaultDBPath = ".sqlite3/nineties.db"
	)

	db, err := db.NewDB(defaultDBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := router.NewRouter(db)

	server := http.Server{
		Addr:    defaultPort,
		Handler: mux,
	}

	errCh := make(chan error)

	// go func() {
	// }
}
