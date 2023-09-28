package main

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	Bff  string `json:"bff"`
}

func main() {
	db := connect()

	createUser(db)
}

func createUser(db *surrealdb.DB) {
	res, err := db.Create("users:bob", User{
		Name: "Bob",
		Bff:  "users:jocke",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func getUsers(db *surrealdb.DB) {
	res, err := db.Query("SELECT * FROM users;", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func connect() *surrealdb.DB {
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": "jocke",
		"pass": "test",
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	return db
}
