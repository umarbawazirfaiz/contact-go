package main

import (
	"contact-go/config"
	"contact-go/handler"
	"contact-go/repository"
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
	contactHandler := handler.NewContactHandler(contactRepo)

	handler.Menu(contactHandler)
}
