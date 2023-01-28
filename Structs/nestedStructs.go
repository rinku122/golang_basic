package main

import "fmt"

type person1 struct {
	firstName string
	lastName  string
	contact   contactDetails
}

type contactDetails struct {
	email string
	phone int // By default zero
}

func makeStructs() {

	alex := person1{firstName: "Rajesh", lastName: "Kumar", contact: contactDetails{email: "rinku@gmail.com", phone: 90111111}}
	fmt.Println(alex) // Bad with structs
	fmt.Printf("%+v", alex)

	// var alex person1

	// alex.firstName = "Rajesh"
	// alex.lastName = "Kumar"
	// alex.contact.email = "rinku"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

}
