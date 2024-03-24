package handler

import (
	"net/http"
	"nineties/service"
)

type CarHandler struct {
	svc *service.CarService
}

func NewCarHandler(svc *service.CarService) *CarHandler {
	return &CarHandler{
		svc: svc,
	}
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
