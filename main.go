package main

import (
	// "encoding/json"
	"fmt"
	// "net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)


type address struct {
	Firstname 	string
	Lastname 	string
	Code 		int
	Phone 		string
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

	addUser(db)

}

func addUser(db *sql.DB) bool {
	statement,_ := db.Prepare("INSERT INTO test.user (ID , FirstName ,LastName ,Age) VALUES (?, ?, ?, ?)")

	defer statement.Close()

	_,err := statement.Exec(1, "Test", "Last" , 25)

	if err != nil {
		panic(err.Error())
		return false
	}
	return true
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