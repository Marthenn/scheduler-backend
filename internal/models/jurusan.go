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
