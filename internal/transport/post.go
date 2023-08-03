package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scheduler-backend/internal/database"
	"scheduler-backend/internal/models"
	"scheduler-backend/internal/services"
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

	res := services.KnapSack(resultJSON.Jurusan, resultJSON.Semester, resultJSON.SKS_Min, resultJSON.SKS_Max)

	fmt.Println(res)

	type body struct {
		Result []models.MataKuliah `json:"MataKuliah"`
		SKS    int                 `json:"SKS"`
		GPA    float64             `json:"GPA"`
	}
	var resBody body
	resBody.Result = res
	if len(res) == 0 {
		resBody.SKS = 0
		resBody.GPA = 0
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"MataKuliah":[],"SKS":0,"GPA":0}`))
		return
	} else {
		resBody.SKS = services.CalculateSKS(res)
		resBody.GPA = services.CalculateGPA(res)
	}

	// return the result
	result, err := json.Marshal(resBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(result))

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
