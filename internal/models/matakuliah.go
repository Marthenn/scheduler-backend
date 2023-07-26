package models

import (
	"github.com/golang-collections/collections/set"
)

var MataKuliahSet = set.New()

type MataKuliah struct {
	ID              int
	Nama            string
	SKS             int
	Jurusan         string
	SemesterMinimal int
	PrediksiNilai   string
}

// NewMataKuliah creates a new MataKuliah object
func NewMataKuliah(ID int, Nama string, SKS int, Jurusan string, SemesterMinimal int, PrediksiNilai string) {
	// create the matakuliah object
	mk := MataKuliah{
		ID:              ID,
		Nama:            Nama,
		SKS:             SKS,
		Jurusan:         Jurusan,
		SemesterMinimal: SemesterMinimal,
		PrediksiNilai:   PrediksiNilai,
	}
	MataKuliahSet.Insert(mk)
}
