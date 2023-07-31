package services

import (
	"fmt"
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

// maxGPA is a helper function for the 1/0 knapsack (return the list with the max GPA).
// If the GPA is the same, return the list with the most mata kuliah
func maxGPA(m1 []models.MataKuliah, m2 []models.MataKuliah) []models.MataKuliah {
	if calculateGPA(m1) > calculateGPA(m2) {
		return m1
	}
	if calculateGPA(m1) == calculateGPA(m2) {
		if len(m1) > len(m2) {
			return m1
		}
	}
	return m2
}

// KnapSack 1/0 will return the best (max GPA) combination of mata kuliah that will satisfy the given constraints (having total sks between minSKS and maxSKS)
// TODO: make the backtrack function, make the minSKS constraint work, and return the list of mata kuliah
func KnapSack(jurusan string, semester int, minSKS int, maxSKS int) []models.MataKuliah {
	// a 2d table to store the dp values
	// dp[i][j] corresponds to GPA with the best combination of mata kuliah for the first i mata kuliah with total sks of j
	dp := make([][]float64, 2) // we only need two rows at a time (space optimization)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, maxSKS+1)
	}
	mk := models.EligibleMatkul(jurusan, semester)

	// filling the dp tables
	for i := 0; i < len(mk); i++ {
		for j := 0; j <= maxSKS; j++ {
			if j-mk[i].SKS >= 0 {
				dp[1][j] = math.Max(dp[0][j], dp[0][j-mk[i].SKS]+indexToGP(mk[i].PrediksiNilai))
			} else {
				dp[1][j] = dp[0][j]
			}
		}

		fmt.Println(dp)

		// move the table up by one row
		for j := 0; j <= maxSKS; j++ {
			dp[0][j] = dp[1][j]
			dp[1][j] = 0
		}
	}

	return []models.MataKuliah{} // TODO: placeholder
}
