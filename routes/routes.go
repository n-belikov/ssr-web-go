package routes

import (
	"github.com/gorilla/mux"
)

type Config struct {
	Pages map[string]string
	Index string
}

var config *Config

func New(cfg *Config) *mux.Router {

	config = cfg
	router := mux.NewRouter()

	router.PathPrefix("/").HandlerFunc(index).Methods("GET")

	return router
}
