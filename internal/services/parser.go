package services

import (
	"encoding/json"
	"scheduler-backend/internal/database"
	"scheduler-backend/internal/models"
)

// ParseJurusan will add one or more jurusan to the database from a .json file
func ParseJurusan(jsonData []byte) {
	var data struct {
		Input []models.Jurusan `json:"Input"`
	}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(data.Input); i++ {
		database.AddJurusan(data.Input[i].Jurusan, data.Input[i].Fakultas)
	}
}

// ParseMatkul will add one or more mata kuliah to the database from a .json file
func ParseMatkul(jsonData []byte) {
	var data struct {
		Input []models.MataKuliah `json:"Input"`
	}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(data.Input); i++ {
		database.AddMatkul(data.Input[i].ID, data.Input[i].Nama, data.Input[i].SKS, data.Input[i].Jurusan, data.Input[i].SemesterMinimal, data.Input[i].PrediksiNilai)
	}
}
