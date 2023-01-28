package main

import "fmt"

type person2 struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	phone int // By default zero
}

func makeStructs1() {

	// alex := person2{firstName: "Rajesh", lastName: "Kumar", contact: contact{email: "rinku@gmail.com", phone: 90111111}}
	// fmt.Println(alex) // Bad with structs
	// fmt.Printf("%+v", alex)

	var alex person2

	alex.firstName = "Rajesh"
	alex.lastName = "Kumar"
	alex.contact.email = "rinku"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

}
