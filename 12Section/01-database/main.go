package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    hashed_password BLOB NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {

	_ = os.Remove("data.db")
	db, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Close database connection")
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created table!!")

	insertQuery := "INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)"

	_, err = db.Exec(insertQuery, "Bharat Bardiya", "bharat@google.com", "this is the way!!!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully inserted values into table!!")

	selectQuery := "SELECT id, name, email, hashed_password, created_at FROM users"

	row, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var name string
		var email string
		var hashedPassword []byte
		var createdAt time.Time
		err = row.Scan(&id, &name, &email, &hashedPassword, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d name: %s email: %s hashedPassword: %s, createdAt: %s\n",
			id, name, email, hashedPassword, createdAt)
	}

}
