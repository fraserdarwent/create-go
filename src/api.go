package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetPerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := Person{Name: "Developer"}
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type Person struct {
	Name string
}
