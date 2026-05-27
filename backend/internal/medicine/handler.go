package medicine

import (
	"encoding/json"
	"net/http"
	"strings"

	"bodhisoft-backend/internal/middleware"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) HandleMedicines(w http.ResponseWriter, r *http.Request) {
	middleware.SetCORSHeaders(w)
	// OPTIONS is a preflight request for CORS, we just need to respond with the appropriate headers and return
	if r.Method == http.MethodOptions {
		return
	}

	switch r.Method {
	case http.MethodGet:
		medicines := h.store.GetAll()
		json.NewEncoder(w).Encode(medicines)

	case http.MethodPost:
		var body CreateMedicineRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if body.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		medicine := h.store.Create(body)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(medicine) // why create a new encoder for each request?

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleMedicineByID(w http.ResponseWriter, r *http.Request) {
	middleware.SetCORSHeaders(w)
	if r.Method == http.MethodOptions {
		return
	}

	//  get id from request url
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/medicines/")

	switch r.Method {
	// case http.MethodGet:
	// 	medicine, err := h.store.GetByID(id)
	// 	if err != nil {
	// 		http.Error(w, "Medicine not found", http.StatusNotFound)
	// 		return
	// 	}

	// 	json.NewEncoder(w).Encode(medicine)
	case http.MethodPut:
		var body UpdateMedicineRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		medicine, err := h.store.Update(id, body)
		if err != nil {
			http.Error(w, "Medicine not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(medicine)

	case http.MethodDelete:
		err := h.store.Delete(id)
		if err != nil {
			http.Error(w, "Medicine not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
