package transport

import (
	"encoding/json"
	"net/http"
	"scheduler-backend/internal/database"
	"scheduler-backend/internal/models"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	database.ReadAll()

	// convert the lists to json
	mataKuliahJSON, err := json.Marshal(models.MataKuliahList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jurusanJSON, err := json.Marshal(models.JurusanList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(mataKuliahJSON)
	w.Write(jurusanJSON)
}
