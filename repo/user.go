package repo

import (
	"database/sql"
	"log"
)

type User struct {
	Id 			int 	`json:"id"`
	Username 	string	`json:"username"`
	Password 	string	`json:"password"`
	Salt 		string	`json:"salt"`
	Email 		string	`json:"email"`
	Active 		bool	`json:"active"`
	Archived 	bool	`json:"archived"`
}

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) FindById(id int) *User {
	u := new(User)
	row := ur.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	
	err := row.Scan(
		&u.Id, 
		&u.Username, 
		&u.Password, 
		&u.Salt, 
		&u.Email, 
		&u.Active, 
		&u.Archived,
	)
	
	if err != nil  {
		if err == sql.ErrNoRows {
			return nil
		} else {
			log.Fatal(err)
		}
	}

	return u
}

func (ur *UserRepository) Save(user *User) int {
	if user == nil {
		return -1
	}

	var insertedId int
	err := ur.DB.QueryRow("INSERT INTO users (username,password,salt,email,active,archived) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id", user.Username, user.Password, user.Salt, user.Email, user.Active, user.Archived).Scan(&insertedId)

	if err != nil {
		log.Fatal(err)
	}

	return insertedId
}