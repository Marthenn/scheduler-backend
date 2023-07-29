package transport

import (
	"fmt"
	"net/http"
)

// postJurusan handles the POST request to /jurusan and adds a new jurusan
func postJurusan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Body)
}

// postMataKuliah handles the POST request to /matakuliah and adds a new mata kuliah
func postMataKuliah(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Body)
}

// postResult handles the POST request to /find and returns the best matkul combination
func postResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Body)
}
