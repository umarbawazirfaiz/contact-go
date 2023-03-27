package model

type ContactRequest struct {
	Name   string `json:"name"`
	NoTelp string `json:"no_telp"`
}

type Contact struct {
	Id     int64
	Name   string
	NoTelp string
}

var Contacts []Contact

// METHOD GET /contacts/ => list contacts
// METHOD POST /contacts/ => insert contact
