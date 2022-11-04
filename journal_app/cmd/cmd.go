package cmd

import (
	"fmt"
	"journal/app"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "entry command",
	Short: "Command to create a new entry",
	Long:  "This command is used to create a new entry",

	Run: func(cmd *cobra.Command, args []string) {
		// check if the user is authorized or not
		fmt.Println(entry, " ", username, " ", password)

		// fmt.Println("Authorization check!")
		// flag = app.ValidateUser(username)

		// if flag {
		//app.Register()

		// instead of add new to map, we should validate the user
		//app.AddNewUserToMap(username)
		//app.DecryptFile()
		app.GetDataFromFile()
		//isValid := app.ValidateUser(username)
		//if isValid {
		fmt.Println("Valid user!!!")
		app.AddNewEntry(username, entry)
		app.AddNewEntryToFile(username)
		app.EncryptFile()
		//} else {
		//fmt.Println("Invalid user!!!")
		//}

		//app.EntryMenu(username)
		//} else {
		//}
	},
}

var entry, username, password, email string
var flag bool

func init() {
	rootCmd.AddCommand(cmd)
	cmd.PersistentFlags().StringVarP(&entry, "add", "a", "", "add a new entry")
	cmd.PersistentFlags().StringVarP(&username, "uname", "u", "", "enter username")
	cmd.PersistentFlags().StringVarP(&password, "passwd", "p", "", "enter password")
	//cmd.PersistentFlags().StringVarP(&email, "email", "p", "", "enter email")

	// required flags
	cmd.MarkPersistentFlagRequired("username")
	cmd.MarkPersistentFlagRequired("password")
}
