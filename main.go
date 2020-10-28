package main

import (
	// "os/user"
	// "encoding/json"
	"fmt"
	// "net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)


type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
	age 		int
}

func main() {
	// handleRequest()
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	

	if err !=  nil {
		fmt.Println("Connect Fail")
	} else {
		fmt.Println("Connect Success")
		defer db.Close()
	}

	// addUser(db)
	fmt.Println(getUser(db))

}

func addUser(db *sql.DB) bool {
	statement,_ := db.Prepare("INSERT INTO test.user (ID , FirstName ,LastName ,Age) VALUES (?, ?, ?, ?)")

	defer statement.Close()

	_,err := statement.Exec(2, "Test", "Last" , 25)

	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

func getUser(db *sql.DB) []User {
	statement,_ := db.Query("select * from test.user")
	defer statement.Close()
	var userList []User
	for statement.Next() {
		var user User
		err := statement.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.age,
		)

		if err != nil {
			panic(err.Error())
		}

		userList = append(userList, user)
	}

	return userList
}

// func homePage(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w,"Welcome to Darkside")
// }

// func handleRequest(){
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/getAddress", getAddress)
// 	http.ListenAndServe(":8080",nil)
// }

// func getAddress(w http.ResponseWriter, r *http.Request){
// 	addBook := address {
// 		Firstname: "Inwza",
// 		Lastname: "007",
// 		Code: 1993,
// 		Phone: "1231234",
// 	}
// 	json.NewEncoder(w).Encode(addBook)
// }