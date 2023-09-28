package main

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
)

func CreateUser(db *surrealdb.DB, id string, user User) {
	res, err := db.Create(id, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func DeleteUser(db *surrealdb.DB, id string) {
	res, err := db.Delete(id)
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

func DeleteCar(db *surrealdb.DB, id string) {
	res, err := db.Delete(id)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
