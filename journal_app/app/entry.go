package app

import (
	"fmt"
	"strings"
	"time"
)

func AddNewEntry(username string, entry string) {
	fmt.Println("Adding a new entry")

	fmt.Println("***username:", userMap)
	getSplitKey(username, entry)

	fmt.Println("Exit from Add new entry function")
}

func getSplitKey(username string, entry string) {
	fmt.Println("Inside getsplitkey function")
	// get the current time
	currentTime := time.Now()
	time := currentTime.Format("03-Nov-22 09:22am")

	for key, val := range userMap {

		fmt.Println("*****uname :", uname, " key : ", key)
		fmt.Println("*****uname :", len(uname), " key : ", len(key))
		splitKey := strings.Split(uname, "---")
		fmt.Println("****SplitKey : ", splitKey[0])
		key = splitKey[0]
		if uname == key {
			fmt.Println("Key found!")
			val.time = append(val.time, time)
			val.data = append(val.data, entry)
			userMap[username] = val
		}
	}
}
