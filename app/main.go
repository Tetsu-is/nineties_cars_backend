package main

import (
	"log"
	"nineties/db"
)

func main() {
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	db, err := db.NewDB(defaultDBPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()

	mux := rest.NewServeMux(db)

}
