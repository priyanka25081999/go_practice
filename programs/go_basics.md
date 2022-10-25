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
