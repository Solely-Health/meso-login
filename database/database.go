package database

import (
	"database/sql"
	"fmt"
	"log"
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

func (database *Database) CloseDB() {
	database.DB.Close()
}
