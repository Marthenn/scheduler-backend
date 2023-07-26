package models

type Jurusan struct {
	Jurusan  string
	Fakultas string
}

// Create a new jurusan object
func NewJurusan(jurusan string, fakultas string) *Jurusan {
	j := Jurusan{
		Jurusan:  jurusan,
		Fakultas: fakultas,
	}
	return &j
}
