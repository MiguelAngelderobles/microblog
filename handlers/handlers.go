package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lutrueba/microblog/middlew"
	"github.com/lutrueba/microblog/routers"
	"github.com/rs/cors"
)

/*Manejadores: seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

}
