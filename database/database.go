package database

import (
	"database/sql"
	"fmt"
	// used for postgres driver
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	DB *sql.DB
}

func (database *Database) InitializeDB(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	// database.DB, err = gorm.Open(Dbdriver, DBURL)
	fmt.Println(DBURL)
	database.DB, err = sql.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the database")
	}
}

func (database *Database) CloseDB() {
	database.DB.Close()
	database.DropTablesDB()
}

func (database *Database) MigrateDB() {
	fmt.Println("Migrating the table")
	createStatement := `
	CREATE TABLE users (
		username varchar(255),
		password varchar(255)
	);
	`
	_, err := database.DB.Exec(createStatement)
	if err != nil {
		panic(err)
	}
}

// DropTablesDB TODO: dont use this in production
func (database *Database) DropTablesDB() {
	fmt.Println("Dropping table")
	createStatement := `
	DROP TABLE users
	`
	_, err := database.DB.Exec(createStatement)
	if err != nil {
		panic(err)
	}
}
