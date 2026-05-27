package main

import (
	"log"
	"net/http"

	"bodhisoft-backend/internal/medicine"
)

func main() {
	medicineStore := medicine.NewStore()
	medicineHandler := medicine.NewHandler(medicineStore)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/v1/medicines", medicineHandler.HandleMedicines)
	http.HandleFunc("/api/v1/medicines/", medicineHandler.HandleMedicineByID)

	log.Println("Server running on :1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}
