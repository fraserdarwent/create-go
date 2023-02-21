package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
