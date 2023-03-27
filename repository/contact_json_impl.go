package repository

import (
	"contact-go/model"
	"encoding/json"
	"errors"
	"os"
)

type contactJsonRepo struct {
}

func NewContactJsonRepository() ContactRepository {
	return &contactJsonRepo{}
}

func (repo *contactJsonRepo) getLastId() int {
	contacts := repo.List()

	tempId := 0
	for _, v := range contacts {
		if tempId < int(v.Id) {
			tempId = int(v.Id)
		}
	}

	return tempId
}

func (repo *contactJsonRepo) getIndexById(id int64) (int, error) {
	contacts := repo.List()
	for i, v := range contacts {
		if v.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("index tidak ditemukan")
}

func (repo *contactJsonRepo) saveToJson(contacts *[]model.Contact) error {
	file, err := os.Create("data/contact.json")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func (repo *contactJsonRepo) List() []model.Contact {
	var contacts []model.Contact

	file, err := os.Open("data/contact.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		panic(err)
	}

	return contacts
}

func (repo *contactJsonRepo) Add(req model.ContactRequest) model.Contact {
	contacts := repo.List()

	// get last id
	lastId := repo.getLastId()

	contact := model.Contact{
		Id:     int64(lastId) + 1,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}

	contacts = append(contacts, contact)

	err := repo.saveToJson(&contacts)
	if err != nil {
		panic(err)
	}

	return contact
}

func (repo *contactJsonRepo) Delete(id int64) error {
	index, err := repo.getIndexById(id)
	if err != nil {
		return err
	}

	contacts := repo.List()

	contacts = append(contacts[:index], contacts[index+1:]...)
	err = repo.saveToJson(&contacts)
	if err != nil {
		panic(err)
	}

	return nil
}
