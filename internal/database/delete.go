package database

import (
	"context"
	"scheduler-backend/internal/models"
)

// ResetDatabase will clear the database
func ResetDatabase() {
	conn := getDBConnection()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "delete from matakuliah")
	if err != nil {
		panic(err)
	}
	models.MataKuliahList = []models.MataKuliah{}
	_, err = conn.Exec(context.Background(), "delete from jurusan")
	if err != nil {
		panic(err)
	}
	models.JurusanList = []models.Jurusan{}
}
