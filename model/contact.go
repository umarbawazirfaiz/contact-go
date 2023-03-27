package model

type ContactRequest struct {
	Name   string
	NoTelp string
}

type Contact struct {
	Id     int64
	Name   string
	NoTelp string
}

var Contacts []Contact

// METHOD GET /contacts/ => list contacts
// METHOD POST /contacts/ => insert contact
