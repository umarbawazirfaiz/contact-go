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

var Contacts []Contact = []Contact{
	{
		Id:     1,
		Name:   "Umar",
		NoTelp: "0812",
	},
}
