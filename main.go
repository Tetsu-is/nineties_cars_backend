package main

import (
	"context"
	"net/http"
	"nineties/db"
	"nineties/handler/router"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

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

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		shutDownWithTimeout(&server, 20)
	case err := <-errCh:
		panic(err)
	}

	stop()
}

func shutDownWithTimeout(srv *http.Server, timeout int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
