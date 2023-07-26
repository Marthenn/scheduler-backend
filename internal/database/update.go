package database

import (
	"context"
	"scheduler-backend/internal/models"
)

/*
Creating an update towards the jurusan will be a bit tricky, because we need to update the mata kuliah as well.
Other than that, jurusan is a primary key, so we can't update it and it will be very unlikely to be updated.
*/

// UpdateMataKuliah will update the mata kuliah with the given ID both in the set and the database
func UpdateMataKuliah(ID string, Nama string, SKS int, Jurusan string, SemesterMinimal int, PrediksiNilai string) {
	if !models.MataKuliahExists(ID) {
		panic("mata kuliah doesn't exist")
	}

	var mk models.MataKuliah // store the old mata kuliah object
	var idx int              // store the index of the mata kuliah object in the list

	for i := 0; i < len(models.MataKuliahList); i++ {
		if models.MataKuliahList[i].ID == ID {
			mk = models.MataKuliahList[i]
			idx = i
			models.MataKuliahList[i].Nama = Nama
			models.MataKuliahList[i].SKS = SKS
			models.MataKuliahList[i].Jurusan = Jurusan
			models.MataKuliahList[i].SemesterMinimal = SemesterMinimal
			models.MataKuliahList[i].PrediksiNilai = PrediksiNilai
			break
		}
	}

	conn := getDBConnection()
	defer conn.Close(context.Background())

	_, err := conn.Exec(context.Background(), "update matakuliah set nama = $1, sks = $2, jurusan = $3, semesterminimal = $4, prediksinilai = $5 where id = $6", Nama, SKS, Jurusan, SemesterMinimal, PrediksiNilai, ID)
	if err != nil {
		models.MataKuliahList[idx] = mk
		panic(err)
	}
}
