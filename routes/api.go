package routes

import (
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)

	text := config.Pages[r.RequestURI]
	w.Write([]byte(strings.Replace(config.Index, "<head>", "<head>"+text, -1)))
}
