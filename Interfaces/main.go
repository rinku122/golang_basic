// package main

// import "fmt"

// type user struct {
// 	name string
// 	age  int
// }

// // type printInterface interface {
// // 	getUserDetails() string
// // }

// func main() {
// 	r := user{name: "rajesh", age: 27}
// 	v := user{name: "vibhuti", age: 36}

// 	fmt.Printf("%+v", r.getUserDetails())
// 	fmt.Printf("%+v", v.getUserDetails())

// }

// func (u user) getUserDetails() user {
// 	// See this as a custom logic
// 	return u
// }

package main

import "fmt"

type userDetails struct{}

type contactDetails struct{}

type printInterface interface {
	getUserDetails() string
}

func main() {
	u := userDetails{}
	c := contactDetails{}

	print(u)
	print(c)
}

func print(p printInterface) {
	fmt.Println(p.getUserDetails())
}

func (u userDetails) getUserDetails() string {
	// See this as a custom logic
	return "Hello"
}

func (c contactDetails) getUserDetails() string {
	// See this as a custom logic
	return "Bye"
}
