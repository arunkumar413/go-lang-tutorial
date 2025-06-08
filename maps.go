package main

import "fmt"

func main() {

	var a = map[string]string{"brand": "Ford", "model": "Mustant", "year": "2025"}
	fmt.Println(a)

	var contacts = map[string]int{"Arun": 9966288507, "Tom": 79833439384}
	fmt.Println((contacts))

	var menu = make(map[string]int)
	menu["Coffee"] = 2
	menu["tea"] = 1
	menu["water bottle"] = 2

}
