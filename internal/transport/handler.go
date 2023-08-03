package transport

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
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
	r.HandleFunc("/", reset).Methods("DELETE")
	r.HandleFunc("/jurusan", getJurusan).Methods("GET")
	r.HandleFunc("/matakuliah", getMatakuliah).Methods("GET")

	handler := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(port, handler)
	if err != nil {
		panic(err)
	}
}
