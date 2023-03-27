package repository

import (
	"contact-go/model"
	"errors"
)

type contactRepository struct {
}

func NewContactRepository() ContactRepository {
	return &contactRepository{}
}

func (repo *contactRepository) getLastId() int {
	contacts := repo.List()

	tempId := 0
	for _, v := range contacts {
		if tempId < int(v.Id) {
			tempId = int(v.Id)
		}
	}

	return tempId
}

func (repo *contactRepository) getIndexById(id int64) (int, error) {
	contacts := repo.List()
	for i, v := range contacts {
		if v.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("index tidak ditemukan")
}

func (repo *contactRepository) List() []model.Contact {
	return model.Contacts
}

func (repo *contactRepository) Add(req model.ContactRequest) model.Contact {
	contact := model.Contact{
		Id:     int64(repo.getLastId()) + 1,
		Name:   req.Name,
		NoTelp: req.NoTelp,
	}
	model.Contacts = append(model.Contacts, contact)

	return contact
}

func (repo *contactRepository) Delete(id int64) error {
	index, err := repo.getIndexById(id)
	if err != nil {
		return err
	}

	model.Contacts = append(model.Contacts[:index], model.Contacts[index+1:]...)
	return nil
}
