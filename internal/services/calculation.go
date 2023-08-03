package services

import (
	"math"
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

// CalculateGPA will calculate the GPA from a list of mata kuliah
func CalculateGPA(matkul []models.MataKuliah) float64 {
	var totalSKS float64
	var totalGP float64

	for _, mk := range matkul {
		totalSKS += float64(mk.SKS)
		totalGP += indexToGP(mk.PrediksiNilai) * float64(mk.SKS)
	}
	return totalGP / totalSKS
}

// CalculateSKS will calculate the total SKS from a list of mata kuliah
func CalculateSKS(matkul []models.MataKuliah) int {
	var totalSKS int
	for _, mk := range matkul {
		totalSKS += mk.SKS
	}
	return totalSKS
}

// bitToMatkul i s a helper function for the 1/0 knapsack (convert a bit to a list of mata kuliah)
func bitToMatkul(bit int, matkul []models.MataKuliah) []models.MataKuliah {
	var ret []models.MataKuliah
	for i := 0; i < len(matkul); i++ {
		if bit&(1<<uint(i)) != 0 {
			ret = append(ret, matkul[i])
		}
	}
	return ret
}

// KnapSack 1/0 will return the best (max GPA) combination of mata kuliah that will satisfy the given constraints (having total sks between minSKS and maxSKS)
func KnapSack(jurusan string, semester int, minSKS int, maxSKS int) []models.MataKuliah {
	mk := models.EligibleMatkul(jurusan, semester)
	table := make([][]float64, 2)
	for i := 0; i < len(table); i++ {
		table[i] = make([]float64, int(math.Pow(2, float64(len(mk)))))
	}

	for i := 0; i < len(mk); i++ {
		increment := 1 << uint(i)                    // the current bit represented as an integer
		upperLimit := int(math.Pow(2, float64(i+1))) // the upper limit of the current mata kuliah

		for j := increment; j < upperLimit; j++ {
			table[1][j] = CalculateGPA(bitToMatkul(j, mk))
		}

		for j := 0; j < len(table[0]); j++ {
			table[0][j] = table[1][j]
		}
	}

	var res []models.MataKuliah

	for i := 0; i < len(table[0]); i++ {
		cur := bitToMatkul(i, mk)
		if CalculateSKS(cur) >= minSKS && CalculateSKS(cur) <= maxSKS {
			if len(res) == 0 || CalculateGPA(cur) > CalculateGPA(res) {
				res = cur
			} else if (CalculateGPA(cur) == CalculateGPA(res)) && (CalculateSKS(cur) > CalculateSKS(res)) {
				res = cur
			}
		}
	}
	return res
}
