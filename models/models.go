package models

import (
	"database/sql"
	"fmt"
	"time"
	
)

const (
	userLimitForNow = 5
	host            = "localhost"
	port            = 5432
	user            = "postgres"
	password        = "postgres"
	dbname          = "postgres"
)

type User struct {
	ID              int
	Name            string
	Email           string
	FollowingMovies []Movie
}

type Movie struct {
	ID             int
	Title          string
	Description    string
	StartYear      int
	EndYear        int
	RuntimeMinutes int
	Genres         string
	NextEpisode    time.Time
}

type DB struct {
	*sql.DB
}

func ResetDB(db *sql.DB) {
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		fmt.Println(err)
	}
}

func (db *DB) Close() {
	db.DB.Close()
}

func (db *DB) Ping() error {
	return db.DB.Ping()
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *DB) Query(query string, args ...interface{}) *sql.Rows {
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}

func (db *DB) UserTableLength() int {
	rows := db.Query("SELECT COUNT(*) FROM users")
	defer rows.Close()
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	return count
}

func (db *DB) ResetDB() {
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		fmt.Println(err)
	}
}

func UserTableLength(db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		fmt.Println(err)
	}

	return count
}

func (db *DB) PrintAllUsers() {
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

func NewDB() *DB {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlconn)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err != nil {
		fmt.Println(err)
	}

	return &DB{db}
}

func (db *DB) CreateUser(name, email string, followingMovies []Movie) {
	_, err := db.Exec("INSERT INTO users (name, email, following_movies) VALUES ($1, $2, $3)", name, email, followingMovies)
	if err != nil {
		fmt.Println(err)
	}
}

func (db *DB) DeleteUser(id int) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	}
}

func (db *DB) UpdateUser(id int, name, email string) {
	_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", name, email, id)
	if err != nil {
		fmt.Println(err)
	}
}
