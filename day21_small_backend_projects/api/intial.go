package api

import (
	"day21/db"
	"day21/model"
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

func (h *Handlers) GetDoughnutsWithType(w http.ResponseWriter, r *http.Request) {
	d_type := r.PathValue("d_type")
	doughnuts, err := h.db.GetDoughnutsWithType(d_type)
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

func (h *Handlers) AddDoughnuts(w http.ResponseWriter, r *http.Request) {
	var doughnuts []model.Doughnut
	if err := json.NewDecoder(r.Body).Decode(&doughnuts); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request...."))
		return
	}
	if err := h.db.AddDoughnuts(doughnuts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Adding got some problem..."))

	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added doughnuts..."))
	return
}
