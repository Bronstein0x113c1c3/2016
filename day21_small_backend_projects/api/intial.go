package api

import (
	"day21/db"
	"encoding/json"
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

func (h *Handlers) GetDoughnuts(w http.ResponseWriter, r *http.Request) {
	doughnuts, err := h.db.GetDoughnuts()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(doughnuts)
	return
}
