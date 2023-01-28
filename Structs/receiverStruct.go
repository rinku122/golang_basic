package main

import "fmt"

// Value types : Use pointer to change this things in a function
//  - int, float, string, bool, structs
// Refrence types : Dont worry about pointer with these things
//	- slices, maps, channels, pointers, fucntions

type person3 struct {
	firstName string
	lastName  string
	contact3
}

type contact3 struct {
	email string
	phone int // By default zero
}

// func makeStructs2() {
// 	alex := person3{firstName: "Rajesh", lastName: "Kumar", contact3: contact3{email: "rinku@gmail.com", phone: 90111111}}
// 	alexpointer := &alex //Gives the address of memory location
// 	alexpointer.updateFirstName("Shkila")
// 	alex.print()
// }

// func (personTopointer *person3) updateFirstName(newName string) {
// 	(*personTopointer).firstName = newName
// }

// func (p person3) print() {
// 	fmt.Printf("%+v", p)

// }

func makeStructs2() {
	alex := person3{firstName: "Rajesh", lastName: "Kumar", contact3: contact3{email: "rinku@gmail.com", phone: 90111111}}
	alex.updateFirstName("Vicky")
	alex.print()
}

func (personTopointer *person3) updateFirstName(newName string) { // In place where there is to be a type like in argument *person3
	//there this means we are telling go that this is a type of pointer
	(*personTopointer).firstName = newName // with pointer if we will use *, it will give us the value
}

func (p person3) print() {
	fmt.Printf("%+v", p)

}
