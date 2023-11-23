package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"ssr-web/routes"
	"time"
)

func main() {
	file := "./config.json"
	template := "./index.html"
	host := "127.0.0.1:8000"

	flag.StringVar(&host, "host", host, "Host")
	flag.StringVar(&file, "file", file, "File config.json")
	flag.StringVar(&template, "template", template, "File index.html")
	flag.Parse()

	index, err := os.ReadFile(template)
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	config := routes.Config{
		Index: string(index),
		Pages: make(map[string]string),
	}

	if err := json.Unmarshal(data, &config.Pages); err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Handler: routes.New(&config),
		Addr:    host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
