package main

import (
	"log"
	"net/http"

	"bodhisoft-backend/internal/db"
	"bodhisoft-backend/internal/medicine"
	"bodhisoft-backend/internal/middleware"
)

func main() {
	db := db.ConnectDB()
	medicineStore := medicine.NewRepo(db)
	medicieneService := medicine.NewService(medicineStore)
	medicineHandler := medicine.NewHandler(medicieneService)

	// testing something here
	db.Exec("SELECT * FROM medicine")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		middleware.SetCORSHeaders(w)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/v1/medicines", medicineHandler.HandleMedicines)
	http.HandleFunc("/api/v1/medicines/", medicineHandler.HandleMedicineByID)

	log.Println("Server running on :1337")
	log.Fatal(http.ListenAndServe(":1337", nil))
}
