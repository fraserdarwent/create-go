package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var router *httprouter.Router

const HOST = "127.0.0.1"
const PORT = "8080"
const ADDRESS = HOST + ":" + PORT

func main() {
	log = Logger()
	router = Router()
	log.Infof("Listening on %v", ADDRESS)
	http.ListenAndServe(ADDRESS, router)
}

func Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/:name", GetPerson)
	return router
}

func Logger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	log = logger.Sugar()
	return log
}
