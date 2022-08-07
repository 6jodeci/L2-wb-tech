package main

import (
	"log"
	"net/http"
	"task11/api"
	"task11/internal/middleware"
)

func main() {
	svc := api.NewEventService()
	h := internal.Logger(internal.New(svc))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	log.Println("Server Started at port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}


}
