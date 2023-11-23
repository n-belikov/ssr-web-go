package routes

import (
	"github.com/gorilla/mux"
)

type Config struct {
	Pages map[string]Page
	Index string
}

type Page struct {
	Head       string `json:"head"`
	StatusCode uint16 `json:"status_code"`
}

var config *Config

func New(cfg *Config) *mux.Router {

	config = cfg
	router := mux.NewRouter()

	router.PathPrefix("/").HandlerFunc(index)

	return router
}
