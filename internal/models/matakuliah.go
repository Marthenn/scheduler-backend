package models

import (
	"github.com/golang-collections/collections/set"
)

var MataKuliahList = []MataKuliah{}
var MataKuliahSet = set.New()

type MataKuliah struct {
	ID              string
	Nama            string
	SKS             int
	Jurusan         string
	SemesterMinimal int
	PrediksiNilai   string
}

// NewMataKuliah creates a new MataKuliah object
func NewMataKuliah(ID string, Nama string, SKS int, Jurusan string, SemesterMinimal int, PrediksiNilai string) {
	// validate each values
	if !validateID(ID) {
		panic("ID must be 2 letters and 4 numbers [A-Z]{2}[0-9]{4}")
	}
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
	MataKuliahList = append(MataKuliahList, mk)
}

// validateID will check if the ID consists of 2letters and 4numbers [A-Z]{2}[0-9]{4}]
func validateID(ID string) bool {
	for i := 0; i < 2; i++ {
		if ID[i] < 'A' || ID[i] > 'Z' {
			return false
		}
	}
	for i := 2; i < 6; i++ {
		if ID[i] < '0' || ID[i] > '9' {
			return false
		}
	}
	return true
}

// validateNilai will check if the PrediksiNilai is within the range of the allowed values
func validateNilai(PrediksiNilai string) bool {
	switch PrediksiNilai {
	case "A", "AB", "B", "BC", "C", "D", "E":
		return true
	default:
		return false
	}
}

// MataKuliahExists will check if a mata kuliah exists in the database from its ID
func MataKuliahExists(ID string) bool {
	for i := 0; i < len(MataKuliahList); i++ {
		if MataKuliahList[i].ID == ID {
			return true
		}
	}
	return false
}

// DeleteMataKuliah will delete a mata kuliah from the set
func DeleteMataKuliah(ID string) {
	if !MataKuliahExists(ID) {
		panic("mata kuliah doesn't exist")
	}
	for i := 0; i < len(MataKuliahList); i++ {
		if MataKuliahList[i].ID == ID {
			MataKuliahList = append(MataKuliahList[:i], MataKuliahList[i+1:]...)
			break
		}
	}
}
