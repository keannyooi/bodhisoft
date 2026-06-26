package medicine

import (
	"encoding/json"
	"net/http"
	"strings"

	"bodhisoft-backend/internal/middleware"
)

type Handler struct {
	service *MedicineService
}

func NewHandler(service *MedicineService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleMedicines(w http.ResponseWriter, r *http.Request) {
	middleware.SetCORSHeaders(w)
	// OPTIONS is a preflight request for CORS, we just need to respond with the appropriate headers and return
	if r.Method == http.MethodOptions {
		return
	}

	switch r.Method {
	case http.MethodGet:
		medicines, err := h.service.GetMedicines()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(medicines)

	case http.MethodPost:
		var body CreateMedicineRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if body.Name == "" || body.Type == "" || body.StrengthUnit == "" {
			http.Error(w, "Name, type, and strength unit are required", http.StatusBadRequest)
			return
		}

		code, err := h.service.CreateMedicine(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(code) // why create a new encoder for each request?

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleMedicineByID(w http.ResponseWriter, r *http.Request) {
	middleware.SetCORSHeaders(w)
	if r.Method == http.MethodOptions {
		return
	}

	//  get code from request url
	code := strings.TrimPrefix(r.URL.Path, "/api/v1/medicines/")

	switch r.Method {
	case http.MethodGet:
		medicine, err := h.service.GetMedicine(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(medicine)
	case http.MethodPut:
		var body UpdateMedicineRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		medicine, err := h.service.UpdateMedicine(code, body)
		if err != nil {
			switch err {
			case ErrMedicineNotFound:
			case ErrInvalidMedicineCode:
				http.Error(w, err.Error(), http.StatusNotFound)
			case ErrInvalidMedicineType:
			case ErrInvalidStrengthUnit:
			case ErrInvalidMedicineStatus:
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(medicine)

	case http.MethodDelete:
		if err := h.service.DeleteMedicine(code); err != nil {
			switch err {
			case ErrMedicineNotFound:
			case ErrInvalidMedicineCode:
				http.Error(w, err.Error(), http.StatusNotFound)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
