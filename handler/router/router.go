package router

import (
	"database/sql"
	"net/http"
	"nineties/handler"
	"nineties/service"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	srv := service.NewCarService(db)

	mux := http.NewServeMux()
	mux.Handle("/hello", handler.NewCarHandler(srv))

	return mux
}
