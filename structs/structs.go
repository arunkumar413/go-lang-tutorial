package main

import "fmt"

func main() {

	// definition of struct

	type Employee struct {
		name        string
		age         int8
		designation string
		location    string
		salary      int32
	}

	var emp1 Employee // create a variable of Employee type

	emp1.name = "arun" // set the properties of struct variable.
	emp1.age = 37
	emp1.designation = "Senior Manager"
	emp1.location = "London"
	emp1.salary = 3433445
	fmt.Println(emp1)
	fmt.Println(emp1.name) // use the dot operator to access the individual properties of a struct.
}
