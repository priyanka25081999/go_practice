**1. How do we run go code in our project?**

  Go to the folder on terminal and run "go run main.go"
	
  1. go run -> take one or two or multiple files, compiles it and executes it.
	
  2. go build -> similar to go run but the difference is that: go run is used for compile and immediately execute the program. while go build is used to just compile a program. It doesn't execute it
	But, this go build command creates an executable file from our source file
	
  3. go fmt -> Format all the code of each file in current directory
	
  4. go install -> compiles and installs a package
	
  5. go get -> Download's the raw/source code of someone else's package
	
  6. go test -> Runs any test associated with the current project


**2. What is package main line ?**

  A package is  a collection of common source code files. Inside go there are two different types of packages: 1. executable and 2. reusable
  
  1. executable: can generate a file which we can actually run (i.e go build command is of executable type)

  2. reusable : code dependacny or libraries. Code use as "helpers". Good place to put reusable logic
	
  3. the word "main" i.e (package main) is used to make an executable type package. It must have main func inside the code
	
  4. if we use any word other than main then it will be create as a reusable/dependancy package and we are not able to execute it




**3. What is import fmt?**

   It means to access some code which is written in another package. import fmt means give my package main, access to all the code and functionality that is contained inside of other package "fmt"
   
   1. fmt is the standard and default go library package. it is shorten form of word format. it uses to print different statement on terminal


**4. what is func means?**

  Similar to other programming languages, func keyword then func name i.e main and then open and close parenthese and then the function body


**5. How is the main.go file organized?**

  At the top we do package declaration, then import other packages that we need, then declare functions (i.e main) to tell golang to do the things
  

**Variables in golang**
1. Go is a statically typed language. Go/C++/Java are statically typed languages while JS/Python/Ruby are dynamically typed languages
	In dynamically typed variables, we do not care the value assigned to the variable. In statically typed language like GO, we assign type of that variable
2. Variables syntax : var card string = "king"
	here, var -> variable starting keyword 2. card -> name of the variable 3. string -> tells the go compiler that only string type values are allowed in this card variable 4. king -> it is the assignment or value holds by that card variable
3. Basic data types in Go -> int, bool, string, float64

          package main

          import "fmt"

          func main() {
            // variables in go
            // two ways of variable declaration

            // 1st: declare and initialize a variable with type
            var card_name1 string = "first card"
            fmt.Println(card_name1)

            // 2nd: declare the variable and type will be identified by the go compiler
            // Note: we use := only at the time of variable declaration.
            // During reassignment of a value we don't use it
            card_name2 := "second card"
            card_name2 = "sec card"
            fmt.Println(card_name2)
          }
4. we can declare the variable outside of any function but cannot assign value in it. we should assign a value inside the function only


**Functions in golang**

new_card() - > func name

string -> return type

func new_card() string {
	// return the new card
	return "third card"
}

        package main

        import "fmt"

        func main() {
          // variables in go
          // two ways of variable declaration

          // 1st: declare and initialize a variable with type
          var card_name1 string = "first card"
          fmt.Println(card_name1)

          // 2nd: declare the variable and type will be identified by the go compiler
          // Note: we use := only at the time of variable declaration.
          // During reassignment of a value we don't use it
          card_name2 := "second card"
          card_name2 = "sec card"
          fmt.Println(card_name2)

          card_name3 := new_card()
          fmt.Println(card_name3)

          fmt.Println(demo_func())
        }

        func new_card() string {
          // return the new card
          return "third card"
        }

        func demo_func() int {
          // return the new card
          return 9
        }

#####
Run function written in another file

        In main.go:

        package main

        func main() {
            printState()
        }


        In a separate file called state.go:

        package main

        import "fmt"

        func printState() {
            fmt.Println("California")
        }

        To run this -> go run main.go state.go
	

**Slices and for loops**

To handle the list of records, go uses 2 data structures : Arrays and Slices
1. Arrays: Fix length of records
2. Slices: An array that can grow or shrink

Every element in an array or slice must be of same type. 
Whenever in go, we see []string or []byte etc then it is the slice of string or slice of byte etc.


	
	
	cards := []string{new_card(), "new_card", new_card()}

	// add new element in a slice
	cards = append(cards, "appended_card")
	fmt.Println(cards)

	// iterate over a slice
	fmt.Println("Printing cards:-")
	for i, card := range cards {
		fmt.Println(i, card)
	}

**OO approach vs GO approach**

1. Go is not an object oriented programming language, so there is no idea of having classes inside of go

2. Using OO approach: we can create a "Deck class" which is a blurprint and can actually contains different methods and properties.
    Then using Deck class we can create multiple instances of a class. 

3. In Go language, we can create a new type from existing primitive data types. main.go->create and manipulate deck, deck.go->code which describes what deck is and how it works, deck_test..go-code to automatically test the deck

**err in go:**

1. It is a value of type "error"

2. If everything goes well, then it will have a value of "nil" (nil means no value), else it will have an actual error message

**go structs:**

1. It is a data structure in golang

2. It is the collection of different properties that are related together

3. definition is depends upon the order of fields defined in the struct. In future, if we change order of fields in the struct then it may create an ambiguity and we need to update the property definition as well.

4. When we don't assign any value to any of the field of struct then go assigns zero value to it

	a. string -> default value -> "" (empty string)

	b. int -> default value -> 0

	c. float -> default value -> 0

	d. bool -> default value -> false

5. Example:
	
		package main
		import "fmt"

		// structure declaration
		type person struct {
			firstName string
			lastName  string
		}

		func main() {
			// first way to pass the property to the struct
			// person1 := person{"Piyu", "Salunke"}

			// second way - pass the property_name:value
			// Even though we change the order of properties in the struct, the program executes correctly.
			// person1 := person{firstName: "Piyu", lastName: "Salunke"}
			// fmt.Println(person1)

			// third way
			var person1 person
			fmt.Println(person1)       // execute successfully, it shows {  } empty strings for struct fields
			fmt.Printf("%+v", person1) // prints all values of struct, o/p as {firstName: lastName:} empty string values

			person1.firstName = "Piyu"
			person1.lastName = "Salunke"
			fmt.Println("\n", person1)
		}
		
**structs in golang:**

1. It is a data structure in golang
2. It is the collection of different properties that are related together
3. definition is depends upon the order of fields defined in the struct. In future, if we change order of fields in the struct then it may create an ambiguity and we need to update the property definition as well.
4. When we don't assign any value to any of the field of struct then go assigns zero value to it
	a. string -> default value -> "" (empty string)
	b. int -> default value -> 0
	c. float -> default value -> 0
	d. bool -> default value -> false
5. Embed different structs, nested or child structs
6. structs with receiver function:
7. Pass by value: Go refers pass by value mechanism. Go copies data and then the copy is available to the another function
8. Pointers in Go
9. Gotchas with Pointers in Go: example:

		func main() {
			mySlice := []string {"Hi","there","how","are","you"}
			mySlice.updateSlice(mySlice);
			fmt.Println(mySlice)
		}

		func updateSlice(s []string) {
			s[0] = "Bye"
		}

	If we run this code, then we can see the value s[0] is updated to Bye. This is a gotcha. Bcz, we have observed that go refers pass by value in structs and how it is changing the value in slice of string. (we are also not using any pointers here).

10. Deep dive into slice:
	Slice is a kind of an array, but it can grow and shrink. Slice interally uses 2 data structures. 1st is slice which has 3 parts: ptr to head, capacity and length. 
	Ptr to head refers to the array which actually stores the elements of a slice, capaity refers to the elements that can contain at present and the length refers to the actual elements present in the slice.
	So, in the memory the 3 parts of slice are stored in one mem location (e.g 0001) and the pt to head array which contains the actual slice elements are stored at another memory location (e.g 0002). 

11. Now, come back to the above code, here Go still creates the copy of the slice (0001 mem location), but the crazy thing is that it is still pointing at the original array in the memory(0002). So we are updating the original array that both copies of slice data structure are now pointing to.

* Along with slice there are couple of other data structures as well which are behaving in the same fashion includes maps, pointers, functions, channels. These are all called as **reference types**. 

* The int, float, string, bool, structs are called as **value types** (use pointers to change the actual values)

**Maps in Golang:**
1. A map is a collection of key-value pair like a dict in python
2. In maps, both key and value are statically typed, they must be of same type.
3. There are different ways to declare a map
4. To access individual values in the map, we need to use square brackets syntax only. i.e colors["red"]. This is because all the keys of the map has types. here we have string. so key should be of string type only and we can define it using square brackets. but, in the struct we use . operator to access the values i.e person1.firstname. here we cannot specify the type i.e string, int etc. In the map, if we use int type then we can define it as colors[10]="xyz". we can easily change it from int to string and viz and also define it and differenciate it

		package main

		import "fmt"

		func main() {
			// Different ways to declare a map
			// first way
			// this is just an empty map, go will initialize it with the default values as we haven't yet initialized it
			// var colors map[string]string
			// fmt.Println(colors)

			// second way - using built-in map function
			// colors := make(map[string]string)
			// colors["red"] = "145"
			// fmt.Println(colors)

			// third way
			colors := map[string]string{
				"red":   "#ff0000",
				"green": "#fff000",
				"white": "#ffffff",
			}

			// iterate over the map
			// here, we are not coping the map, we are passing the reference here. Bcz map is a reference type data structure
			printMap(colors)

			// delete the key
			// syntax : delete(map_name, key_name)
			delete(colors, "green")
			fmt.Println("After delete :", colors)
		}

		func printMap(m map[string]string) {
			for color, hex := range m {
				fmt.Println("Hex code for", color, "is", hex)
			}
		}


**Interfaces in Golang:**
1. Golang doesn't support function overloading

		package main

		import "fmt"

		type bot interface {
			getGreeting() string
		}

		type englishBot struct{}

		type spanishBot struct{}

		func main() {
			eb := englishBot{}
			sb := spanishBot{}

			printGreeting(eb)
			printGreeting(sb)
		}

		func printGreeting(b bot) {
			fmt.Println(b.getGreeting())
		}

		func (eb englishBot) getGreeting() string {
			return "Hi there!"
		}

		func (sb spanishBot) getGreeting() string {
			return "Hola!"
		}
2. Interfaces are not generic types
