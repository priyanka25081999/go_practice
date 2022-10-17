// REST API : GET, POST

package main

// import the packages
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// To create a REST API, we are using gin library
// Install the dependancy using go mod init example/folder-name (go.mod)
// To gett the dependanices locally use: go get github.com/gin-gonic/gin (go.sum)

// define the structure
type todo struct {
	Id			string	`json:"id"` 	// use of string literal so that go will understand
	Name 		string	`json:"name"`
	Completed	bool	`json:"completed"`
}

var todos = []todo {
	{Id: "1", Name: "Task1", Completed: true},
	{Id: "2", Name: "Task2", Completed: false},
	{Id: "3", Name: "Task3", Completed: false},
}

// context contains bunch of info coming from HTTP request
// context will also convert the request data into json (added json in struct)
func get_todos(context *gin.Context) {
	// first para is the http status
	// second para is the json which is the resp. It will auto convert data into json
	context.IndentedJSON(http.StatusOK, todos)
}

func add_todo(context *gin.Context) {
	var newTodo todo 

	// BindJson get the json passed by client in the put request body
	// and convert it into the go understandable structure i.e todo
	// this method will return an error if the structure doesn't match
	// so, we are catching an error in "err" variable below
	if err := context.BindJSON(&newTodo);
	err != nil {
		// if nil means no error and everything is ok, then bind the json to todo struct
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	// create a server
	router := gin.Default()

	// create an endpoint
	router.GET("/todos", get_todos)	
	router.POST("/todos", add_todo)
	router.Run("localhost:9090")
}

// To run : use http://localhost:9090 on your browser. 
// Get method : You can see output as https://github.com/priyanka25081999/go_practice/blob/main/output/get_method.JPG
// POST method : Send request as https://github.com/priyanka25081999/go_practice/blob/main/output/post%20method.JPG
