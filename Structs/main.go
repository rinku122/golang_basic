package main

type person struct {
	firstName string
	lastName  string
}

func main() {

	// alex := person{firstName: "Rajesh", lastName: "Kumar"}
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	var alex person

	alex.firstName = "Rajesh"
	alex.lastName = "Kumar"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// makeStructs()
	// makeStructs1()
	makeStructs2()

}
