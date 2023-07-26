package database

import (
	"context"
	"scheduler-backend/internal/models"
)

// DeleteJurusan will delete a jurusan from the database if and only if the jurusan doesn't have any mata kuliah
func DeleteJurusan(jurusan string) {
	if !models.JurusanExists(jurusan) {
		panic("jurusan doesn't exist")
	}
	if models.HasMataKuliah(jurusan) {
		panic("jurusan still has mata kuliah")
	}

	conn := getDBConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "delete from jurusan where jurusan = $1", jurusan)
	if err != nil {
		panic(err)
	}

	models.DeleteJurusan(jurusan)
}

// DeleteMataKuliah will delete a mata kuliah from the database
func DeleteMataKuliah(ID string) {
	if !models.MataKuliahExists(ID) {
		panic("mata kuliah doesn't exist")
	}

	conn := getDBConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "delete from matakuliah where id = $1", ID)
	if err != nil {
		panic(err)
	}

	models.DeleteMataKuliah(ID)
}
