package main

import (
	"buyback-service/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error Panic: ", err)
		}
	}()

	r := mux.NewRouter()

	routes.Routes(r)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
