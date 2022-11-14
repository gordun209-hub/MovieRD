package main

import (
	_ "github.com/lib/pq"
)

// simdilk dble urasmayak
func main() {
	Main()
	db := NewDB()
	defer db.Close()
}
