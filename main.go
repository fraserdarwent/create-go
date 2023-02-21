package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

const HOST = "127.0.0.1"
const PORT = "8080"
const ADDRESS = HOST + ":" + PORT

func main() {
	logger, _ := zap.NewProduction()
	log = logger.Sugar()

	router := httprouter.New()
	router.GET("/api/v1/:name", GetName)
	log.Infof("Listening on %v", ADDRESS)
	http.ListenAndServe(ADDRESS, router)
}

func GetName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
