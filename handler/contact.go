package handler

import (
	"contact-go/helper"
	"contact-go/model"
	"contact-go/repository"
	"fmt"
)

type contactHandler struct {
	ContactRepository repository.ContactRepository
}

func NewContactHandler(contactRepository repository.ContactRepository) ContactHandler {
	return &contactHandler{contactRepository}
}

func (handler *contactHandler) List() {
	helper.CallClear()
	fmt.Printf("ID\t\t| Nama\t\t| No Telp\t\t\n")
	contacts := handler.ContactRepository.List()
	for _, v := range contacts {
		fmt.Printf("%d\t\t| %s\t\t| %s\t\t\n", v.Id, v.Name, v.NoTelp)
	}
	var back string
	fmt.Scanln(&back)
	Menu(handler)
}

func (handler *contactHandler) Add() {
	helper.CallClear()
	fmt.Println("Add new contact")
	fmt.Print("Name = ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("No Telp = ")
	var no_telp string
	fmt.Scanln(&no_telp)

	contactRequest := model.ContactRequest{
		Name:   name,
		NoTelp: no_telp,
	}

	contact := handler.ContactRepository.Add(contactRequest)
	if contact.Id != 0 {
		fmt.Println("Berhasil add contact with id", contact.Id)
	}
	var back string
	fmt.Scanln(&back)
	Menu(handler)
}

func (handler *contactHandler) Update() {
}

func (handler *contactHandler) Delete() {
	helper.CallClear()
	fmt.Println("Delete contact")
	fmt.Print("ID = ")
	var id int64
	fmt.Scanln(&id)

	err := handler.ContactRepository.Delete(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	var back string
	fmt.Scanln(&back)
	Menu(handler)
}
