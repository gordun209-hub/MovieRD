package main

import (
	"context"
	"fmt"
	"os"

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
	ID    int
	Name  string
	Email string
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

	var greeting string

	err := conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	// display database

	// Insert Users table
	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %va", err)
	}

	conn.InsertUser("laaa", "mww")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert user: %v", err)
	}

	// print users
	conn.getUsers()
}

func (db *DB) getUsers() {
	rows := db.Query("SELECT * FROM users")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(id, name, email)
	}
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

	return User{id, name, email}
}
