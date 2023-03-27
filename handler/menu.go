package handler

import (
	"contact-go/helper"
	"fmt"
)

func Menu(handler ContactHandler) {
	helper.CallClear()
	fmt.Println("Menu")
	fmt.Println("1. list contact")
	fmt.Println("2. add contact")
	fmt.Println("3. delete contact")

	var menu int
	fmt.Print("Pilih menu ")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		handler.List()
	case 2:
		handler.Add()
	case 3:
		handler.Delete()
	}
}
