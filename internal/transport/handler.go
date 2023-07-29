package transport

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// REST is the main function that will handle all the REST requests, it also handles the CORS
func REST() {
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

	r.HandleFunc("/matakuliah", postMataKuliah).Methods("POST")
	r.HandleFunc("/jurusan", postJurusan).Methods("POST")
	r.HandleFunc("/find", postResult).Methods("POST")
	r.HandleFunc("/reset", reset).Methods("DELETE")
	r.HandleFunc("/", getAll).Methods("GET")

	handler := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
