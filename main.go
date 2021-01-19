package main

import (
	"log"

	"github.com/MiguelAngelderobles/microblog/bd"
	"github.com/MiguelAngelderobles/microblog/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
