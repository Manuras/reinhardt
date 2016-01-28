package repo

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	// Move this to a test database helper
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	DB_TEST_USER = "postgres"
	DB_TEST_PASSWORD = "root"
	DB_TEST_NAME = "reinhardt_test"
)

// Integration test for User repo
func TestUser(t *testing.T) {
	Convey("User", t, func() {
		dbString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_TEST_USER, DB_TEST_PASSWORD, DB_TEST_NAME)
		db, err := sql.Open("postgres", dbString)

		if err != nil {
			panic(err)
		}

		repo := UserRepository{db}

		Convey("Save", func() {
			repo.Save(CreateUser())

			var count int 
			db.QueryRow("select count(*) from users").Scan(&count)
		
			So(count, ShouldEqual, 1)
		})

		Convey("FindById", func() {
			created_user := CreateUser()
			id := repo.Save(created_user)
			user := repo.FindById(id)

			So(id, ShouldEqual, user.Id)
			So(created_user.Username, ShouldEqual, user.Username)
		})

		Reset(func() {
			db.Query("TRUNCATE TABLE users")
		})
	})
}

func CreateUser() *User {
	user := new(User)
	user.Username = "Test"
	user.Password = "password"
	user.Salt = "salt"
	user.Email = "test@test.nl"
	user.Active = true
	user.Archived = false

	return user
}