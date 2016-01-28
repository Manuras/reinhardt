package main

import (
	"github.com/manuras/reinhardt/repo"

	"database/sql"
	_ "github.com/lib/pq"

	//"encoding/json"
	"fmt"
	"log"
)

type Repositories struct {
	user *repo.UserRepository
}

const (
	DB_USER = "postgres"
	DB_PASSWORD = "root"
	DB_NAME = "reinhardt"
)

func main() {

	// Load config

	// Create logger

	// Setup database
	dbString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbString)

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	// Setup all repositories
	var repos = Repositories{&repo.UserRepository{db}}

	user := new(repo.User)
	user.Username = "test"

	repos.user.Save(user)
}
