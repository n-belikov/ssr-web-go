package routes

import (
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	page, exists := config.Pages[r.URL.Path]

	if exists {
		w.WriteHeader(int(page.StatusCode))
		w.Write([]byte(strings.Replace(config.Index, "<head>", "<head>"+page.Head, -1)))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(config.Index))
	}
}
