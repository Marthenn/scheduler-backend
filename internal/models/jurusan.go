package models

var JurusanList = []Jurusan{}

type Jurusan struct {
	Jurusan  string
	Fakultas string
}

// NewJurusan creates a new jurusan object and adds it to the set
func NewJurusan(jurusan string, fakultas string) {
	j := Jurusan{
		Jurusan:  jurusan,
		Fakultas: fakultas,
	}
	JurusanList = append(JurusanList, j)
}

// JurusanExists check if a jurusan exists in the set (the fakultas will be ignored)
func JurusanExists(jurusan string) bool {
	for i := 0; i < len(JurusanList); i++ {
		if JurusanList[i].Jurusan == jurusan {
			return true
		}
	}
	return false
}

// HasMataKuliah check if a jurusan has any mata kuliah
func HasMataKuliah(jurusan string) bool {
	for i := 0; i < len(MataKuliahList); i++ {
		if MataKuliahList[i].Jurusan == jurusan {
			return true
		}
	}
	return false
}

// DeleteJurusan will remove a jurusan from the set
func DeleteJurusan(jurusan string) {
	if !JurusanExists(jurusan) {
		panic("jurusan doesn't exist")
	}
	for i := 0; i < len(JurusanList); i++ {
		if JurusanList[i].Jurusan == jurusan {
			JurusanList = append(JurusanList[:i], JurusanList[i+1:]...)
			break
		}
	}
}
