package app

import (
	"scheduler-backend/internal/database"
	"scheduler-backend/internal/transport"
)

// Run is the entrypoint of the application (the main function)
func Run() {
	database.ReadJurusan()
	database.ReadMataKuliah()
	transport.REST()
}
