package main

import (
	// "os/user"
	// "encoding/json"
	"fmt"
	// "net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"helloworld/route"
	
)


type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
	age 		int
}

func main() {
	
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	

	if err !=  nil {
		fmt.Println("Connect Fail")
	} else {
		fmt.Println("Connect Success")
		defer db.Close()
	}
	fmt.Println(getUser(db))
	route.HandleRequest()

	// addUser(db)
	// updateUser(db,1)
	

}


func addUser(db *sql.DB) bool {
	statement,_ := db.Prepare("INSERT INTO test.person (FirstName ,LastName ,Age) VALUES (?, ?, ?)")

	defer statement.Close()

	_,err := statement.Exec("Test3", "Last" , 35)

	if err != nil {
		panic(err.Error())
	}
	return true
}

func getUser(db *sql.DB) []User {
	result,_ := db.Query("select * from test.person")
	defer result.Close()
	var userList []User
	for result.Next() {
		var user User
		err := result.Scan(
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

func updateUser(db *sql.DB ,id int) bool {
	statement,_ := db.Prepare("update test.person set  FirstName = ? ,LastName = ? ,Age = ? where id = ?")

	defer statement.Close()

	_,err := statement.Exec("Test", "Last" , 25, id)

	if err != nil {
		panic(err.Error())
	}
	return true
}

func deleteUser(db *sql.DB, id int) bool {
	statement,_ := db.Prepare("delete from test.person where id = ?")
	defer statement.Close()
	_,fail := statement.Exec(id)

	if fail != nil {
		panic(fail.Error())
	}
	return true
}

//----------------route -----------------

// func homePage(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w,"Welcome to Darkside")
// }

// func handleRequest(){
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/getAddress", getAddress)
// 	http.ListenAndServe(":8080",nil)
// }

// func getAddress(w http.ResponseWriter, r *http.Request){
// 	addBook := User {
// 		FirstName: "Inwza",
// 		LastName: "007",
// 		age: 25,
// 	}
// 	json.NewEncoder(w).Encode(addBook)
// }