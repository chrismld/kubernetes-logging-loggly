package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)

		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   10,
		}).Info("A group of walrus emerges from the ocean")

		log.WithFields(log.Fields{
			"omg":    true,
			"number": 122,
		}).Warn("The group's number increased tremendously!")

		fmt.Fprintf(w, "Hello World!\n")
	})

	http.ListenAndServe(":8080", nil)
}
