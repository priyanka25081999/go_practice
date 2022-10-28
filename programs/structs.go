package main

import "fmt"

// structure declaration
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// first way to pass the property to the struct
	// person1 := person{"Piyu", "Salunke"}

	// second way - pass the property_name:value
	// Even though we change the order of properties in the struct, the program executes correctly.
	// person1 := person{firstName: "Piyu", lastName: "Salunke"}
	// fmt.Println(person1)

	// third way
	// var person1 person
	// fmt.Println(person1)       // execute successfully, it shows {  } empty strings for struct fields
	// fmt.Printf("%+v", person1) // prints all values of struct, o/p as {firstName: lastName:} empty string values

	// person1.firstName = "Piyu"
	// person1.lastName = "Salunke"
	// fmt.Println("\n", person1)

	// Nested struct
	person1 := person{
		firstName: "Piyu",
		lastName:  "Salunke",
		contact: contactInfo{
			email:   "salunkeps99@gmail.com",
			zipCode: 411028,
		},
	}

	fmt.Println(person1)
}
