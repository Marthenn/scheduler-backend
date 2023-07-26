package models

import (
	"github.com/golang-collections/collections/set"
)

var JurusanSet = set.New()

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
	JurusanSet.Insert(j)
}

// JurusanExists check if a jurusan exists in the set (the fakultas will be ignored)
func JurusanExists(jurusan string) bool {
	res := false
	JurusanSet.Do(func(j interface{}) {
		if j.(Jurusan).Jurusan == jurusan {
			res = true
			return
		}
	})
	return res
}

// HasMataKuliah check if a jurusan has any mata kuliah
func HasMataKuliah(jurusan string) bool {
	res := false
	MataKuliahSet.Do(func(mk interface{}) {
		if mk.(MataKuliah).Jurusan == jurusan {
			res = true
			return
		}
	})
	return res
}

// DeleteJurusan will remove a jurusan from the set
func DeleteJurusan(jurusan string) {
	var fakultas string
	if !JurusanExists(jurusan) {
		panic("jurusan doesn't exist")
	}
	JurusanSet.Do(func(j interface{}) {
		if j.(Jurusan).Jurusan == jurusan {
			fakultas = j.(Jurusan).Fakultas
			return
		}
	})
	JurusanSet.Remove(Jurusan{Jurusan: jurusan, Fakultas: fakultas})
}
