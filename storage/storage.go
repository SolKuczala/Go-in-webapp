package storage

import (
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Storage struct {
	UsersTable []*User
}

func (s *Storage) Connect() {
	db, err := sql.Open("mysql", "b86f0a6107db68:2a1d431a@eu-cdbr-west-03.cleardb.net/heroku_e02a84e34d46887?reconnect=true")
	if err != nil {
		panic(err.Error())
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	log.Fatal("CONECTATED!!!!")
	defer db.Close()
}

func (s *Storage) SaveNewUser(user *User) error {
	/*
		//create new user (sign up using password)
		INSERT INTO user_auth(email, pass, google_token)
		VALUES ("pepe@pepe","123456789", NULL);

		//create new user (sign up using google)
		INSERT INTO user_auth(email, pass, google_token)
		VALUES ("pepe@pepe",NULL, "a231d3asd3ass132a46s5d46sad");
	*/

	//Generate an ID for the user
	user.ID = 1
	//Append new user to alluser Array and Save it
	s.UsersTable = append(s.UsersTable, user)
	log.Printf("Storage saved new user: %+v\n", user)
	log.Printf("%+v\n", s.UsersTable)
	return nil
}

func (s *Storage) GetUserAuth(user *User) error {
	/*
		// get user auth for login
		SELECT id, email
		FROM user_auth
		WHERE email="input email" AND pass="input password";
	*/
	user.Email = "pepe@papa"
	user.ID = 1
	log.Printf("Storage found login data: %+v\n", user)
	return nil
}

func (s *Storage) GetUserProfile(user *User) error {
	/*
		SELECT full_name, address, phone
		FROM user_info
		WHERE id={EL ID DEL LOGIN};
	*/
	return nil
}
