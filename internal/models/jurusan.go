package models

import (
	"github.com/golang-collections/collections/set"
)

var jurusanSet = set.New()

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
	jurusanSet.Insert(j)
}
