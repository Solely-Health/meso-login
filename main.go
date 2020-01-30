package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	// "github.com/meso-org/meso-login/database"
)

var (
	db Database
)

type Database struct {
	DB *sql.DB
}

func (database *Database) InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	// database.DB, err = gorm.Open(Dbdriver, DBURL)
	database.DB, err = sql.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}
}

// User asdasd
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
	// this is database vars
	var (
		DbHost     = "datastore"
		DbUser     = "steven"
		DbPassword = "password"
		DbNam      = "meso_datastore"
		DbPort     = "5432"
	)
	db.InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbNam)

	router := mux.NewRouter()
	router.HandleFunc("/register", registerUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
