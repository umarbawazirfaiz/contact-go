package repository

import "contact-go/model"

type ContactRepository interface {
	List() []model.Contact
	Add(req model.ContactRequest) model.Contact
	Delete(id int64) error
}
