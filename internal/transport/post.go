package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scheduler-backend/internal/database"
	"scheduler-backend/internal/models"
)

// postJurusan handles the POST request to /jurusan and adds a new jurusan
func postJurusan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	type data struct {
		Input []models.Jurusan `json:"Input"`
	}
	var jurusanJSON data
	err := json.NewDecoder(r.Body).Decode(&jurusanJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, jurusan := range jurusanJSON.Input {
		database.AddJurusan(jurusan.Jurusan, jurusan.Fakultas)
	}
}

// postMataKuliah handles the POST request to /matakuliah and adds a new mata kuliah
func postMataKuliah(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	type data struct {
		Input []models.MataKuliah `json:"Input"`
	}
	var mataKuliahJSON data
	err := json.NewDecoder(r.Body).Decode(&mataKuliahJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, mataKuliah := range mataKuliahJSON.Input {
		database.AddMatkul(mataKuliah.ID, mataKuliah.Nama, mataKuliah.SKS, mataKuliah.Jurusan, mataKuliah.SemesterMinimal, mataKuliah.PrediksiNilai)
	}
}

// postResult handles the POST request to /find and returns the best matkul combination
func postResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Body)
}
