package main

import (
	"badminton-go/db"
	"badminton-go/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)

	http.HandleFunc("/match/list", service.MatchListHandler)
	http.HandleFunc("/match/create", service.MatchCreateHandler)
	http.HandleFunc("/match/join", service.MatchJoinHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
