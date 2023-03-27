package main

import (
	"contact-go/handler"
	"contact-go/repository"
)

func main() {
	contactRepo := repository.NewContactRepository()
	contactHandler := handler.NewContactHandler(contactRepo)

	handler.Menu(contactHandler)
}
