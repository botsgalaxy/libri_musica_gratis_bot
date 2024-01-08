package modules

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type User struct {
	UserId    int64
	FirstName string
	LastName  string
	Username  string
	CreatedAt time.Time
}

func (u *User) SaveUser() error {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	statement, err := db.Prepare(`
			INSERT INTO users (user_id, first_name, last_name, username, created_at)
			VALUES (?, ?, ?, ?, ?)
		`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(u.UserId, u.FirstName, u.LastName, u.Username, u.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Println("New User saved successfully")
	return nil
}

func getUsers()([]User, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT user_id, first_name, last_name, username, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Username, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}