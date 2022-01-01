package main

import (
	"fmt"
	"log"
	"net/http"
	"badminton-go/db"
	"badminton-go/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}