package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kondratenko-go-dev/go-architecture/internal/person"
)

type personRequest struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	City   string `json:"city"`
	Street string `json:"street"`
}

type personResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	City   string `json:"city"`
	Street string `json:"street"`
}

func main() {
	storage := person.NewMemoryStorage()
	service := person.NewService(storage)

	r := chi.NewRouter()

	r.Get("/people", func(w http.ResponseWriter, r *http.Request) {
		handleListPeople(w, r, service)
	})

	r.Post("/people", func(w http.ResponseWriter, r *http.Request) {
		handleCreatePerson(w, r, service)
	})

	r.Get("/people/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleGetPerson(w, r, service)
	})

	r.Delete("/people/{id}", func(w http.ResponseWriter, r *http.Request) {
		handleDeletePerson(w, r, service)
	})

	log.Println("Starting server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}

func handleListPeople(w http.ResponseWriter, r *http.Request, service *person.Service) {
	people, err := service.ListPeople()
	if err != nil {
		respondError(w, http.StatusInternalServerError, "failed to list people")
		return
	}

	resp := make([]personResponse, 0, len(people))
	for _, p := range people {
		resp = append(resp, personResponse{
			ID:     p.ID(),
			Name:   p.Name(),
			Age:    p.Age(),
			City:   p.City(),
			Street: p.Street(),
		})
	}
	respondJSON(w, http.StatusOK, resp)
}

func handleCreatePerson(w http.ResponseWriter, r *http.Request, service *person.Service) {
	var req personRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	p, err := service.CreatePerson(req.Name, req.Age, req.City, req.Street)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := personResponse{
		ID:     p.ID(),
		Name:   p.Name(),
		Age:    p.Age(),
		City:   p.City(),
		Street: p.Street(),
	}

	respondJSON(w, http.StatusCreated, resp)
}

func handleGetPerson(w http.ResponseWriter, r *http.Request, service *person.Service) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	p, err := service.GetPerson(id)
	if err != nil {
		if errors.Is(err, person.ErrNotFound) {
			respondError(w, http.StatusNotFound, "person not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "failed to get person")
		return
	}

	resp := personResponse{
		ID:     p.ID(),
		Name:   p.Name(),
		Age:    p.Age(),
		City:   p.City(),
		Street: p.Street(),
	}

	respondJSON(w, http.StatusOK, resp)
}

func handleDeletePerson(w http.ResponseWriter, r *http.Request, service *person.Service) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = service.DeletePerson(id)
	if err != nil {
		if errors.Is(err, person.ErrNotFound) {
			respondError(w, http.StatusNotFound, "person not found")
			return
		}
		respondError(w, http.StatusInternalServerError, "failed to delete person")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("failed to write response:", err)
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{
		"error": message,
	})
}
