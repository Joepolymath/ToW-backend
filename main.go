package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	WriteTimeout = 90
	ReadTimeout = 90
)

func main() {
	r := mux.NewRouter()
	srv := http.Server{
		Handler: r,
		Addr: "0.0.0.0:8001",
		WriteTimeout: WriteTimeout * time.Second,
		ReadTimeout: ReadTimeout * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}