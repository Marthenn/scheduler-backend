package services

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	fmt.Println(data.Input)
	fmt.Println(reflect.TypeOf(data.Input))
}
