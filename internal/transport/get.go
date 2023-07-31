package transport

import (
	"encoding/json"
	"fmt"
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

// getResult handles the POST request to /find and returns the best matkul combination
func getResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// read the request body
	type data struct {
		Jurusan  string `json:"Jurusan"`
		Semester int    `json:"Semester"`
		SKS_Min  int    `json:"SKS_Min"`
		SKS_Max  int    `json:"SKS_Max"`
	}
	var resultJSON data
	err := json.NewDecoder(r.Body).Decode(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(resultJSON.Jurusan)
	fmt.Println(resultJSON.Semester)
	fmt.Println(resultJSON.SKS_Min)
	fmt.Println(resultJSON.SKS_Max)

	res := models.MataKuliahList

	// return the result
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
