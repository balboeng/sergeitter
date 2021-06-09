package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/balboeng/sergeitter/mw"
	"github.com/balboeng/sergeitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers set port and listen the server*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", mw.CheckDbStatus(routers.Register))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("<<--- PORT: " + PORT + " ------->>")
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
