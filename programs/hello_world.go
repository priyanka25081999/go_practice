// Packages are very important in go language
// when we specify same pkg name in multiple files
// Then go compiler will append all files in one main file which is called main
// which will contain the code of all different files

package main

// import fmt package for formatting i.e printing statements or logging
import "fmt"

// Every main pkg can have the main function, this is the entry point for the go code
func main() {
	fmt.Println("Hello world...!")
}

/* Output
To run the file : go run <file_name>
C:\Golang\Programs>go run hello_world.go
Hello world...!

To build the project : go build <file_name>
C:\Golang\Programs>go build hello_world.go

This will execute and create a binary (.exe) file into same folder
Then we can do ./<file_name> to execute it and get the output
C:\Golang\Programs>hello_world.exe
Hello world...!
*/
