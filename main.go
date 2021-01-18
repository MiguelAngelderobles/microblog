package main

import (
	"log"

	"github.com/MiguelAngelderobles/microblog/handlers"
)

func main() {
	if database.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
