package service

import "database/sql"

type CarService struct {
	db *sql.DB
}

func NewCarService(db *sql.DB) *CarService {
	return &CarService{
		db: db,
	}
}

// func (srv *CarService) CreateCar(ctx context.Context, car *Car) error {}
