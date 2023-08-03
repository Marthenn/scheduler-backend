package transport

import (
	"encoding/json"
	"net/http"
	"scheduler-backend/internal/models"
)

// getJurusan will return all the jurusan from the database
func getJurusan(w http.ResponseWriter, r *http.Request) {
	// convert the list to json
	jurusanJSON, err := json.Marshal(models.JurusanList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jurusanJSON)
}

// getMataKuliah will return all the mata kuliah from the database
func getMatakuliah(w http.ResponseWriter, r *http.Request) {
	// convert the list to json
	matakuliahJSON, err := json.Marshal(models.MataKuliahList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(matakuliahJSON)
}
