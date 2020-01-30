package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(w, "U dun fugged up")
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Fprintf(w, "U dun fugged up")
		fmt.Println(err)
	}
	fmt.Println(user.Username, user.Password)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", registerUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
