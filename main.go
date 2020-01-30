package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meso-org/meso-login/database"
)

var (
	db database.Database
)

// User structsaf asds
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
	// saveUserToDB(user)
	fmt.Println(user.Username, user.Password)
}

// func saveUserToDB(userInfo User) {
// 	sqlStatement := `
// 	INSERT INTO users_table (username, password)
// 	VALUES ($1, $2)`
// 	_, err := db.Exec(sqlStatement, userInfo.Username, userInfo.Password)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {

	// TODO: use env package to grab env values
	var (
		DbHost     = "datastore"
		DbUser     = "steven"
		DbPassword = "password"
		DbName     = "meso_datastore"
		DbPort     = "5432"
	)
	db.InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbName)
	db.MigrateDB()
	defer db.CloseDB()

	router := mux.NewRouter()
	router.HandleFunc("/register", registerUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
