package main

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

func CreateUser(db *surrealdb.DB) {
	res, err := db.Create("users:bob", User{
		Name: "Bob",
		Bff:  "users:jocke",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func DeleteUser(db *surrealdb.DB) {
	res, err := db.Delete("users:bob")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func GetUsers(db *surrealdb.DB) {
	res, err := db.Query("SELECT * FROM users;", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func CreateCar(db *surrealdb.DB, id string, car Car) {
	res, err := db.Create(id, car)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
