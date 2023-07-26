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
	// validate each values
	if SKS < 1 {
		panic("sks must be greater than 0")
	}
	if SemesterMinimal < 1 || SemesterMinimal > 8 {
		panic("semester minimal must be between 1 and 8")
	}
	if !validateNilai(PrediksiNilai) {
		panic("prediksi nilai must be one of A, AB, B, BC, C, D, E")
	}

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

// validateNilai will checks if the PrediksiNilai is within the range of the allowed values
func validateNilai(PrediksiNilai string) bool {
	switch PrediksiNilai {
	case "A", "AB", "B", "BC", "C", "D", "E":
		return true
	default:
		return false
	}
}
