package storage

import (
	"errors"
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"password"`
	Gtoken   string `json:"access_token"`
}

type Storage struct {
	db *sql.DB
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) Init() {
	db, err := sql.Open("mysql", "b07ba4b7aa60e9:4b7b92b3@tcp(eu-cdbr-west-03.cleardb.net:3306)/heroku_7e1b871d7963fd5")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database is up!")
	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	s.db = db
}

/*SaveNewUser Receives the user input to be posted to the DB
 */
func (s *Storage) SaveNewUser(user *User) error {
	if user.Password == "" {
		return errors.New("Missing password field")
	}
	q := `insert into user_auth(email, pass, google_token) VALUES (?,?,?);`
	_, err := s.db.Query(q, user.Email, user.Password, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SaveNewUserFromGoogleAuth(user *User) error {
	if user.Gtoken == "" {
		return errors.New("Missing token field")
	}
	q := `insert into user_auth(email, pass, google_token) VALUES (?,?,?);`
	_, err := s.db.Query(q, user.Email, nil, user.Gtoken)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUserAuth(user *User) error {
	/*
		// get user auth for login
		SELECT id, email
		FROM user_auth
		WHERE email="input email" AND pass="input password";
	*/

	// how to extract the results
	//for rows.Next() {
	//var (
	//id   int64
	//name string
	//)
	//if err := rows.Scan(&id, &name); err != nil {
	//log.Fatal(err)
	//}
	//log.Printf("id %d name is %s\n", id, name)
	//}
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
