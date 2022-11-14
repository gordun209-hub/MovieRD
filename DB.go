package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	userLimitForNow = 5
	host            = "localhost"
	port            = 5432
	user            = "postgres"
	password        = "postgres"
	dbname          = "postgres"
)

type DB struct {
	*pgx.Conn
}

type User struct {
	ID              int
	Name            string
	Email           string
	FollowingMovies []Movie
}

type Movie struct {
	ID              int
	Name            string
	NextReleaseDate time.Time
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email)
}

func (u User) UpdateName(name string) {
	// TODO
}

func (u User) UpdateEmail(email string) {
	// TODO
}

func (u User) Delete() {
	// TODO
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}

func NewDB() *DB {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	conn, err := pgx.Connect(context.Background(), psqlconn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
		os.Exit(1)
	}
	return &DB{conn}
}

func (db *DB) Close() {
	db.Conn.Close(context.Background())
}

func (db *DB) Query(query string) pgx.Rows {
	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v", err)
		os.Exit(1)
	}
	return rows
}

func Main() {
	conn := NewDB()
	defer conn.Close()

	_, err := conn.Exec(
		context.Background(),
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %va", err)
	}
}

func (db *DB) getUsers() []User {
	rows := db.Query("SELECT * FROM users")
	defer rows.Close()
	var users []User
	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err, "err")
		}

		users = append(users, User{id, name, email, nil})
	}

	return users
}

func (db *DB) InsertUser(name string, email string) {
	if db.ChechIfUserExists(name) {
		return
	}

	_, err := db.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		fmt.Println(err)
	}
}

func (db *DB) ClearTable() {
	_, err := db.Exec(context.Background(), "DROP TABLE IF EXISTS users")
	if err != nil {
		fmt.Println(err)
	}
}

func (db *DB) ChechIfUserExists(name string) bool {
	var count int
	err := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE name = $1", name).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}

	return count > 0
}

func (db *DB) GetUser(name string) User {
	var id int
	var email string
	err := db.QueryRow(context.Background(), "SELECT * FROM users WHERE name = $1", name).Scan(&id, &name, &email)
	if err != nil {
		fmt.Println(err)
	}

	return User{id, name, email, nil}
}

func (db *DB) GetAllUsers() []User {
	rows := db.Query("SELECT * FROM users")
	defer rows.Close()
	var users []User
	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err)
		}

		users = append(users, User{id, name, email, nil})
	}

	return users
}
