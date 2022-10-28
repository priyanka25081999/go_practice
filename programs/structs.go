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

	person1.print()

	// Give us the access to memory address where the struct person1 is sitting at.
	// Here person1Pointer is not directly refer to the struct (person1)
	// Instead, it's a reference to the memory address that struct exists at.
	person1Pointer := &person1
	person1Pointer.updateName("Priyanka")

	// second way to call the updateName function:
	// person1.updateName("Priyanka")
	// No need to create person1Pointer. Golang will automatically convert and allow us to call the function
	// and treat as pass by reference

	person1.print()
}

// receiver function
func (p person) print() {
	fmt.Println(p)
}

// receiver function : update the property (firstname) of the struct
// Note : This will not update the actual value of the firstname, you can still see "Piyu" as firstname
// This is due to pass by value mechanism used by Golang
// func (p person) updateName(name string) {
// 	p.firstName = name
// }

// Pointers : updatename receiver function with pointer
// pass by reference
// *person receiver is the type which we are interested
func (pointerToPerson *person) updateName(name string) {
	// *pointerToPerson (operator)- take the memory address and give us the value that exists at that memory address.
	(*pointerToPerson).firstName = name
}
