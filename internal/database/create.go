package database

import (
	"context"
	"scheduler-backend/internal/models"
)

// AddJurusan adds a new jurusan to the database
func AddJurusan(jurusan string, fakultas string) {
	conn := getDBConnection()
	defer conn.Close(context.Background())

	// create the jurusan object
	models.NewJurusan(jurusan, fakultas)

	_, err := conn.Exec(context.Background(), "insert into jurusan values ($1, $2)", jurusan, fakultas)
	if err != nil {
		models.JurusanSet.Remove(models.Jurusan{Jurusan: jurusan, Fakultas: fakultas})
		panic(err)
	}
}

// AddMatkul adds a new matkul to the database
func AddMatkul(ID string, Nama string, SKS int, Jurusan string, SemesterMinimal int, PrediksiNilai string) {
	// check if the jurusan exists, if not ask to add the jurusan first
	if !models.JurusanExists(Jurusan) {
		panic("jurusan doesn't exist")
	}

	// create the matkul object
	models.NewMataKuliah(ID, Nama, SKS, Jurusan, SemesterMinimal, PrediksiNilai)

	conn := getDBConnection()
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), "insert into matakuliah values ($1, $2, $3, $4, $5, $6)", ID, Nama, SKS, Jurusan, SemesterMinimal, PrediksiNilai)
	if err != nil {
		models.MataKuliahSet.Remove(models.MataKuliah{ID: ID, Nama: Nama, SKS: SKS, Jurusan: Jurusan, SemesterMinimal: SemesterMinimal, PrediksiNilai: PrediksiNilai})
		panic(err)
	}
}
