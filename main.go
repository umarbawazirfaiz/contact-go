package main

import (
	"contact-go/config"
	"contact-go/handler"
	"contact-go/repository"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	var contactRepo repository.ContactRepository
	switch config.Storage {
	case "json":
		contactRepo = repository.NewContactJsonRepository()
	default:
		contactRepo = repository.NewContactRepository()
	}
	// contactHandler := handler.NewContactHandler(contactRepo)
	// handler.Menu(contactHandler)

	contactHandlerHttp := handler.NewContactHandlerHttp(contactRepo)
	NewServer(contactHandlerHttp)
}

func NewServer(handler handler.ContactHandlerHttp) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler.List(w, r)
		} else if r.Method == "POST" {
			handler.Add(w, r)
		}
	})

	server := http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
