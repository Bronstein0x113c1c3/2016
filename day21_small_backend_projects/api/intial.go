package api

import (
	"day21/db"
	"net/http"
)

type Handlers struct {
	db db.IDoughnutDatabase
}

func InitHandlers(db db.IDoughnutDatabase) *Handlers {
	return &Handlers{
		db: db,
	}
}

func (h *Handlers) GetDoughnuts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
