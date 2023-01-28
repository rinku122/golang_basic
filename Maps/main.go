package main

// func main() {

// 	// var colors map[string]string for only initializing

// 	colors := map[string]string{
// 		"red":   "#kkkkkk",
// 		"black": "#000000",
// 	}

// 	colors["red"] = "#10000"

// 	fmt.Println(colors)

// 	delete(colors, "red")

// 	fmt.Println(colors)

// }

// func main() {

// 	colors := make(map[int]string)

// 	colors[10] = "#10000"
// 	colors[20] = "#20000"

// 	fmt.Println(colors)

// 	delete(colors, 10)

// 	fmt.Println(colors)

// }

func main() {

	// var colors map[string]string for only initializing

	colors := map[string]string{
		"red":    "#kkkkkk",
		"black":  "#000000",
		"whitte": "#fffff",
	}

	printAll(colors)

}

func printAll(c map[string]string) {
	for color, hex := range c {
		println("Hex code for", color, "is", hex)
	}
}
