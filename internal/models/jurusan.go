package models

type Jurusan struct {
	Jurusan  string
	Fakultas string
}

// NewJurusan creates a new jurusan object
func NewJurusan(jurusan string, fakultas string) *Jurusan {
	j := Jurusan{
		Jurusan:  jurusan,
		Fakultas: fakultas,
	}
	return &j
}
