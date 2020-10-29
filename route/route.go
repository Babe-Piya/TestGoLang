package route

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID 			int
	FirstName 	string
	LastName 	string
	age 		int
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome to Darkside")
}

func HandleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddress)
	http.ListenAndServe(":8080",nil)
}

func getAddress(w http.ResponseWriter, r *http.Request){
	addBook := User {
		FirstName: "Inwza",
		LastName: "007",
		age: 25,
	}
	json.NewEncoder(w).Encode(addBook)
}
