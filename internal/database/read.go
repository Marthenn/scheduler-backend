package database

import (
	"context"
	"scheduler-backend/internal/models"
)

// ReadJurusan reads all jurusan from the database
func ReadJurusan() {
	models.JurusanList = []models.Jurusan{} // reset the list

	conn := getDBConnection()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from jurusan")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var jurusan string
		var fakultas string
		err = rows.Scan(&jurusan, &fakultas)
		if err != nil {
			panic(err)
		}
		models.NewJurusan(jurusan, fakultas)
	}
}

// ReadMataKuliah reads all mata kuliah from the database
func ReadMataKuliah() {
	models.MataKuliahList = []models.MataKuliah{} // reset the list

	conn := getDBConnection()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select * from matakuliah")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var ID string
		var Nama string
		var SKS int
		var Jurusan string
		var SemesterMinimal int
		var PrediksiNilai string
		err = rows.Scan(&ID, &Nama, &SKS, &Jurusan, &SemesterMinimal, &PrediksiNilai)
		if err != nil {
			panic(err)
		}
		models.NewMataKuliah(ID, Nama, SKS, Jurusan, SemesterMinimal, PrediksiNilai)
	}
}

// ReadAll reads all data from the database
func ReadAll() {
	ReadJurusan()
	ReadMataKuliah()
}
