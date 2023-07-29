package transport

import (
	"net/http"
	"scheduler-backend/internal/database"
)

// reset will reset the database
func reset(w http.ResponseWriter, r *http.Request) {
	database.ResetDatabase()
	w.WriteHeader(http.StatusOK)
}
