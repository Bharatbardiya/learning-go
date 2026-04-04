package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL, 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreateAt       time.Time `json:"created_at"`
}

func main() {

	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connection established")

	krishna, err := getUserbyEmail(db, "krishna@dummy.com")
	if err != nil {
		log.Fatal(err)
	}

	user_json, err := json.MarshalIndent(krishna, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(user_json))

	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	users_json, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(users_json))
}

func getUserbyEmail(db *sql.DB, email string) (*User, error) {
	var usr User
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`
	row := db.QueryRow(stmt, email)

	err := row.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.HashedPassword, &usr.CreateAt)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	var users []User
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreateAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
