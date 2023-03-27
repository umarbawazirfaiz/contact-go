package handler

import (
	"contact-go/model"
	"contact-go/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type contactHandlerHttp struct {
	contactRepository repository.ContactRepository
}

func NewContactHandlerHttp(contactRepository repository.ContactRepository) ContactHandlerHttp {
	return &contactHandlerHttp{contactRepository}
}

func (handler *contactHandlerHttp) List(w http.ResponseWriter, r *http.Request) {
	contacts := handler.contactRepository.List()

	byteContacts, err := json.Marshal(&contacts)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	w.Write(byteContacts)
}

func (handler *contactHandlerHttp) Add(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request body error"))
	}

	req := model.ContactRequest{}
	decode := json.NewDecoder(r.Body)
	decode.Decode(&req)

	contact := handler.contactRepository.Add(req)
	byteContact, err := json.Marshal(contact)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(byteContact)
}
