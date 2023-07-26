package services

import (
	"scheduler-backend/internal/models"
)

// indexToGP will convert the index to the GP according to ITB's standard
func indexToGP(index string) float64 {
	switch index {
	case "A":
		return 4.0
	case "AB":
		return 3.5
	case "B":
		return 3.0
	case "BC":
		return 2.5
	case "C":
		return 2.0
	case "D":
		return 1.0
	case "E":
		return 0.0
	default:
		panic("invalid index")
	}
}

// calculateGPA will calculate the GPA from a list of mata kuliah
func calculateGPA(matkul []models.MataKuliah) float64 {
	var totalSKS float64
	var totalGP float64

	for _, mk := range matkul {
		totalSKS += float64(mk.SKS)
		totalGP += indexToGP(mk.PrediksiNilai) * float64(mk.SKS)
	}
	return totalGP / totalSKS
}

// KnapSack 1/0 will return the best (max GPA) combination of mata kuliah that will satisfy the given constraints
func KnapSack(jurusan string, semester int, minSKS int, maxSKS int, mk []models.MataKuliah) []models.MataKuliah {

}
