package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	CreationDate string `json:"date"`
}

var newUser User

func CreateUser(context *gin.Context) {
	//var newUser User

	err := context.BindJSON(&newUser)
	if err != nil {
		fmt.Println("Error occured during CreateUser", err)
		return
	}

	fmt.Println("User details:")
	fmt.Println("username:", newUser.Username, "\nname:", newUser.Name, "\nemail:", newUser.Email)

	// add data into database
	db, err := sql.Open("mysql", "root:root123@tcp(127.17.0.1:3306)/userdb")
	if err != nil {
		fmt.Println("Error during connection establishment!", err)
	}

	_, err = db.Exec("insert into users VALUES ('" + newUser.Username + "', '" + newUser.Username + "', '" + newUser.Username + "')")

	if err != nil {
		fmt.Println("Sorry, unable to insert data into table, Please try again", err)
	}

	defer db.Close()

	fmt.Println("Data added successfully!")
	context.IndentedJSON(http.StatusCreated, newUser)
}

func GetUser(context *gin.Context) {

}

func GetAllUsers(context *gin.Context) {
	// TODO: get the data from database and store it into newUser and return newUser

	context.IndentedJSON(http.StatusOK, newUser)
}

func DeleteUser(context *gin.Context) {

}

func GetUserMeta(context *gin.Context) {
	// get the username first
	// uname := context.Param("username")

}
