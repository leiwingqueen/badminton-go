package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Resp struct {
	Code int
	Msg  string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	fmt.Printf("name:%s\n", name)
	resp := Resp{0, "success"}
	s, _ := json.Marshal(resp)
	fmt.Fprintln(w, string(s))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/list", ListHandler)
	http.ListenAndServe(":8000", nil)
}
